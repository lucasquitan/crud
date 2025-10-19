package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasquitan/crud/src/configuration/logger"
	"github.com/lucasquitan/crud/src/controller/routes"
)

var (
	PORT = "PORT"
)

func main() {
	logger.Info("🚀 Server is running on port " + PORT)

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Info("Error loading .env file")
	}

	if PORT == "" {
		PORT = "8080"
	}
	
	// Initialize router
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
