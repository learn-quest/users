package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn-quest/users/models"
	"github.com/learn-quest/users/services"
)

func Singup(c *gin.Context) {
	// creating empty user to bind it with req body
	var user models.User
	// binding req body to user
	c.ShouldBindBodyWithJSON(&user)

	// checking if any of below fields are missed in req body
	if user.Name == "" || user.Email == "" || user.Username == "" || user.Country == "" {
		// returning bad req error if any field is missing
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Please provide all fields in request body",
		})
		return
	}
	// calling service to insert user in db
	err := services.InserUser(c, &user)
	if err != nil {
		// if any error returned by service while inserting data into database
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Something went wrong",
			"data": err.Error(),
		})
		return
	}
	// returning 201 if user created successfully
	c.JSON(http.StatusCreated, gin.H{
		"msg":  "Signed up successfully",
		"data": user,
	})
}
