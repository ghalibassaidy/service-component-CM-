package handler

import (
	"net/http"
	"service_components/internal/database"
	"service_components/internal/model"
	"service_components/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
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

func CreateTag(c *gin.Context) {
	var input CreateTagRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	slug := strings.ToLower(strings.ReplaceAll(input.Name, " ", "-"))

	tag := model.Tag{
		Name: input.Name,
		Slug: slug,
	}

	if err := database.DB.Create(&tag).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Gagal menyimpan tag ke database")
		return
	}

	utils.Created(c, tag)
}

func GetAllTags(c *gin.Context) {
	var tags []model.Tag

	if err := database.DB.Order("name asc").Find(&tags).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Gagal mengambil data tag")
		return
	}

	utils.Success(c, tags)
}
