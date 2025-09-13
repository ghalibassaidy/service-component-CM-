package database

import (
	"fmt"
	"log"
	"service_components/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	var err error

	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("FATAL: Failed to connectioni database: %v", err)
	}

	fmt.Println("Connection Succes")
}
