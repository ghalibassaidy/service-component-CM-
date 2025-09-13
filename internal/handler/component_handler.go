package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"service_components/internal/database"
	"service_components/internal/model"
	"service_components/internal/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddComponentTagRequest struct {
	TagID uuid.UUID `json:"tag_id" binding:"required"`
}

type UpdateComponentRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CreateComponentRequest struct {
	Name            string      `json:"name" binding:"required"`
	Description     string      `json:"description"`
	CategoryID      uuid.UUID   `json:"category_id" binding:"required"`
	CodeJSX         string      `json:"code_jsx" binding:"required"`
	CodeCSS         string      `json:"code_css"`
	PropsDefinition interface{} `json:"props_definition"`
}

// CreateCategory godoc
// @Summary Membuat kategori baru
// @Description Endpoint untuk menambah kategori baru
// @Tags Category
// @Accept json
// @Produce json
// @Param data body CreateCategoryRequest true "Data kategori"
// @Success 201 {object} model.Category
// @Failure 400 {object} utils.ErrorResponse
// @Router /categories [post]

func CreateComponent(c *gin.Context) {
	var input CreateComponentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var propsJSON []byte
	var err error
	if input.PropsDefinition != nil {
		propsJSON, err = json.Marshal(input.PropsDefinition)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Gagal mengkonversi props_definition")
			return
		}
	}

	slug := strings.ToLower(strings.ReplaceAll(input.Name, " ", "-"))

	component := model.Component{
		Slug:            slug,
		Name:            input.Name,
		Description:     input.Description,
		CategoryID:      input.CategoryID,
		CodeJSX:         input.CodeJSX,
		CodeCSS:         input.CodeCSS,
		PropsDefinition: propsJSON,
		UserID:          uuid.New(),
	}

	if err := database.DB.Create(&component).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create component: "+err.Error())
		return
	}
	var createdComponent model.Component
	if err := database.DB.Preload("Category").Preload("Tags").First(&createdComponent, "id = ?", component.ID).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Gagal mengambil data yang baru dibuat")
		return
	}

	utils.Created(c, createdComponent)
}

func GetAllComponents(c *gin.Context) {
	var components []model.Component
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	err := database.DB.Preload("Category").Preload("Tags").
		Offset(offset).Limit(limit).Order("created_at desc").
		Find(&components).Error

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch components")
		return
	}

	utils.Success(c, components)
}

func GetComponentBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var component model.Component

	err := database.DB.Preload("Category").Preload("Tags").Where("slug = ?", slug).First(&component).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, "Component Not Found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch component")
		return
	}

	utils.Success(c, component)
}

func UpdateComponentBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var component model.Component

	err := database.DB.Where("slug = ?", slug).First(&component).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, "Component Not Found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to query component")
		return
	}

	var input UpdateComponentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Name != nil {
		component.Name = *input.Name
		component.Slug = strings.ToLower(strings.ReplaceAll(*input.Name, " ", "-"))
	}
	if input.Description != nil {
		component.Description = *input.Description
	}

	if err := database.DB.Save(&component).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to update component")
		return
	}

	utils.Success(c, component)
}

func DeleteComponentBySlug(c *gin.Context) {
	slug := c.Param("slug")

	result := database.DB.Where("slug = ?", slug).Delete(&model.Component{})

	if result.Error != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to delete component")
		return
	}

	if result.RowsAffected == 0 {
		utils.Error(c, http.StatusNotFound, "Component Not Found")
		return
	}

	c.Status(http.StatusNoContent)
}

func AddComponentTag(c *gin.Context) {
	componentSlug := c.Param("slug")

	var input AddComponentTagRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var component model.Component
	err := database.DB.Where("slug = ?", componentSlug).First(&component).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, "Component Not Found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to find component")
		return
	}

	var tag model.Tag
	err = database.DB.First(&tag, input.TagID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Error(c, http.StatusNotFound, "Tag Not Found")
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to find tag")
		return
	}

	if err := database.DB.Model(&component).Association("Tags").Append(&tag); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to add tag to component")
		return
	}

	database.DB.Preload("Category").Preload("Tags").First(&component)

	utils.Success(c, component)
}
