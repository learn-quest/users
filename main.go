package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	dbSetup "github.com/learn-quest/users/config"
	"github.com/learn-quest/users/middlewares"
	"github.com/learn-quest/users/routes"
)

func main() {
	// getting environment from os env to read appropriate env file from env folder
	environment := os.Getenv("environment")
	if environment != "" && environment == "production" {
		// To check if environment present and its production
		godotenv.Load("./env/.env.production")
		fmt.Println("Application starting :: mode=" + environment)
	} else {
		// for development environment
		godotenv.Load("./env/.env.development")
		fmt.Println("Application starting :: mode=development")
	}
	session := dbSetup.InitDBConnection()
	defer session.Close()

	// creating parent router
	router := gin.Default()
	api := router.Group("/api")
	{
		// injecting session in middlewares
		api.Use(middlewares.DbSession(session))
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from learn quest users service",
			})
		})
		routes.MainRouter(api)
	}
	PORT := os.Getenv("PORT")
	router.Run(":" + PORT)
}
