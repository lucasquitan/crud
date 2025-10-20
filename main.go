package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasquitan/crud/src/configuration/logger"
	"github.com/lucasquitan/crud/src/controller"
	"github.com/lucasquitan/crud/src/controller/routes"
	"github.com/lucasquitan/crud/src/model/service"
)

func main() {
	logger.Info("🚀 Server is running")

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Info("Error loading .env file")
	}

	// Initialize router
	router := gin.Default()

	// Initialize dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	routes.InitRoutes(&router.RouterGroup, userController)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
