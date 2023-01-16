package database

import (
	"devbook-api/pkg/database/migrations"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect is responsible for connecting to the database
func Connect() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal("Error: ", err)
	}

	port, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		port = 5432
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DATABASE"), port)

	fmt.Println("Connecting to the database")
	fmt.Println("DSN: ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the database")
		log.Fatal("Error: ", err)
	}

	return db, nil
}

// StartDB is responsible for opening the database connection
func StartDB() (*gorm.DB, error) {
	db, _ := Connect()

	_ = migrations.Migrate(db)

	return db, nil
}

// GetDB is responsible for getting the database connection
func GetDB() *gorm.DB {
	return db
}

// Close is responsible for closing the database connection
func Close(db *gorm.DB) error {
	d, err := db.DB()

	if err != nil {
		fmt.Println("Could not get database")
		log.Fatal("Error: ", err)
	}

	d.Close()

	return nil
}
