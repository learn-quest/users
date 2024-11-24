package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// getting environment from os env to read appropriate env file from env folder
	environment := os.Getenv("environment")
	if environment != "" && environment == "production" {
		// To check if environment present and its production
		godotenv.Load("./env/production.env")
		fmt.Println("Application starting :: mode=" + environment)
	} else {
		// for development environment
		godotenv.Load("./env/development.env")
		fmt.Println("Application starting :: mode=development")
	}
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from learn quest users service",
			})
		})
	}
	PORT := os.Getenv("PORT")
	r.Run(":" + PORT)
}
