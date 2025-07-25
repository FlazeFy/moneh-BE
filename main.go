package main

import (
	"log"
	"moneh/config"
	"moneh/models"
	"moneh/modules"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func initLogging() {
	now := time.Now()
	logFileName := "logs/moneh-" + now.Format("January-2006") + ".log"

	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		panic("Error loading ENV")
	}

	// Connect DB
	db := config.ConnectDatabase()
	MigrateAll(db)

	// Setup Gin
	router := gin.Default()
	redisClient := config.InitRedis()

	modules.SetUpDependency(router, db, redisClient)

	// Run server
	port := os.Getenv("PORT")
	router.Run(":" + port)

	log.Printf("Moneh is running on port %s\n", port)
}

func MigrateAll(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Error{},
		&models.Dictionary{},
		&models.User{},
		&models.Admin{},
		&models.History{},
		&models.Feedback{},
		&models.Pocket{},
		&models.Flow{},
		&models.FlowRelation{},
	)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Migrate Success!")
}
