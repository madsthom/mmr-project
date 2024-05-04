package database

import (
	"example.com/m/v2/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    // Define the database configuration
    db, err := gorm.Open(sqlite.Open("./db/foosball.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // Auto Migrate your models
    err = db.AutoMigrate(&models.User{}, &models.Team{}, &models.Match{}) 
    if err != nil {
        panic("failed to auto migrate models")
    }

    // Assign the database connection to the global variable
    DB = db
}
