package database

import (
	//"example.com/m/v2/db/models" // Import your models package
	"gorm.io/driver/mysql"       // Import the MySQL driver
	"gorm.io/gorm"               // Import GORM
)

var DB *gorm.DB

func InitDatabase() {
    // Define the database configuration
    dsn := "user:password@tcp(your-database-url:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

    // Open a connection to the database
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // // Auto Migrate your models
    // err = db.AutoMigrate(&models.User{}, &models.Team{}, &models.Match{}) 
    // if err != nil {
    //     panic("failed to auto migrate models")
    // }

    // Assign the database connection to the global variable
    DB = db
}
