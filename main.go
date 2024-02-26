package main

import (
	"fmt"
	"log"

	"dealls/config"
	"dealls/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	// Initialize Gin router
	router := gin.Default()

	routes.SetupRoutes(router)

	// Start the server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
