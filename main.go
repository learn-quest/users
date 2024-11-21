package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hi from Main")
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from learn quest from users",
			})
		})
	}
	r.Run(":5000")
}
