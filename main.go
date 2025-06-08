package main

import (
	"fmt"
	"moneh/config"
	"moneh/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		panic("error loading ENV")
	}

	config.ConnectDatabase()

	// Connect DB
	db := config.ConnectDatabase()
	MigrateAll(db)

	// Setup Gin
	router := gin.Default()

	// Run server
	router.Run(":9000")
}

func MigrateAll(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Migrate Success!")
}
