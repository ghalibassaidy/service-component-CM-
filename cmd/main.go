package main

import (
	"log"

	"service_components/internal/config"
	"service_components/internal/database"
	"service_components/internal/handler"
	"service_components/internal/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "service_components/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ComponentHub API
// @version 1.0
// @description This is the API for the ComponentHub marketplace.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)
	database.Seeder()
	err := database.DB.AutoMigrate(&model.Category{}, &model.Tag{}, &model.Component{})
	if err != nil {
		log.Fatalf("FATAL: Failed to migrate database: %v", err)
	}
	
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")
	{
		api.GET("/health", handler.HealthCheck)
		api.POST("/components", handler.CreateComponent)
		api.GET("/components", handler.GetAllComponents)
		api.GET("/components/:slug", handler.GetComponentBySlug)
		api.PATCH("/components/:slug", handler.UpdateComponentBySlug)
		api.DELETE("/components/:slug", handler.DeleteComponentBySlug)
		api.POST("/components/:slug/tags", handler.AddComponentTag)

		api.POST("/categories", handler.CreateCategory)
		api.GET("/categories", handler.GetAllCategories)

		api.PATCH("/components/:slug/status", handler.UpdateComponentStatus)
		api.PATCH("/components/:slug/approval", handler.UpdateComponentApproval)

		api.POST("/tags", handler.CreateTag)
		api.GET("/tags", handler.GetAllTags)
	}

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
