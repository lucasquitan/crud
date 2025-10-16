package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasquitan/crud/src/controller/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("🚀 Server is running on port %s\n", os.Getenv("PORT"))

	// Initialize router
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
	
}
