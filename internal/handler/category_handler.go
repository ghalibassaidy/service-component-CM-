package handler

import (
	"net/http"
	"service_components/internal/database"
	"service_components/internal/model"
	"service_components/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
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

func CreateCategory(c *gin.Context) {
	var input CreateCategoryRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	slug := strings.ToLower(strings.ReplaceAll(input.Name, " ", "-"))

	category := model.Category{
		Name: input.Name,
		Slug: slug,
	}
	if err := database.DB.Create(&category).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Gagal menyimpan kategori ke database")
		return
	}

	utils.Created(c, category)
}

func GetAllCategories(c *gin.Context) {
	var categories []model.Category

	if err := database.DB.Order("name asc").Find(&categories).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Gagal mengambil data kategori")
		return
	}

	utils.Success(c, categories)
}
