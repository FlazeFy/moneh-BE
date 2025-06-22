package main

import (
	"fmt"
	"log"
	"moneh/config"
	"moneh/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func initLogging() {
	f, err := os.OpenFile("pelita.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	initLogging()
	log.Println("Moneh service is starting...")

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
	port := os.Getenv("PORT")
	router.Run(":" + port)

	log.Printf("Pelita is running on port %s\n", port)
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
