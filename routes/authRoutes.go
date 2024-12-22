package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learn-quest/users/controllers"
)

func AuthRouter(api *gin.RouterGroup) {
	authRouter := api.Group("/auth")
	{
		authRouter.GET("/sign-up", controllers.Singup)
	}
}
