package main

import (
	"fmt"
	"log"

	"dealls/config"
	"dealls/database"
	"dealls/repository"
	"dealls/routes"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	db = database.Connect()

	userRepository := repository.NewUserRepository(db)
	premiumRepository := repository.NewPremiumRepository(db)

	// Initialize Gin router
	router := gin.Default()

	routes.SetupRoutes(router, userRepository, premiumRepository)

	// Start the server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	//defer db.Close()
}
