package migrations

import (
	"devbook-api/pkg/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Migrate is responsible for running the migrations
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println("Could not migrate")
		log.Fatal("Error: ", err)
	}

	return nil
}
