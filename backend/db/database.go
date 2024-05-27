package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mmr/backend/db/models"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Define the database configuration
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	migrateErr := db.AutoMigrate(&models.User{}, &models.Team{}, &models.Match{}, &models.PlayerHistory{}, &models.MMRCalculation{})
	if migrateErr != nil {
		panic("failed to migrate database")
	}
	// Assign the database connection to the global variable
	DB = db
}
