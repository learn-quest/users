package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn-quest/users/models"
	"github.com/learn-quest/users/services"
)

func Singup(c *gin.Context) {
	var user models.User
	c.ShouldBindBodyWithJSON(&user)
	if user.Name == "" || user.Email == "" || user.Username == "" || user.Country == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Please provide all fields in request body",
		})
		return
	}
	err := services.InserUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Something went wrong",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Testing signup api from controller",
	})
}
