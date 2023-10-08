package database

import (
	"Assignment-3/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     = "localhost"
	dbPort     = 5434
	dbUser     = "postgres"
	dbPassword = "gendut167"
	dbName     = "Assignment-3"
	db         *gorm.DB
)

func StartDB() {
	Config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error

	// Initialize the database using GORM.
	db, err = gorm.Open(postgres.Open(Config), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate model ke database
	db.AutoMigrate(&models.SensorData{})
	fmt.Println("Connected to the database")
}

// GetDB mengembalikan instance DB
func GetDB() *gorm.DB {
	return db
}
