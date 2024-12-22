package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/learn-quest/users/controllers"
)

func AuthRouter(api *gin.RouterGroup) {
	// Auth routes: /api/auth
	authRouter := api.Group("/auth")
	{
		// Signup route: /api/auth/sign-up
		authRouter.GET("/sign-up", controllers.Singup)
	}
}
