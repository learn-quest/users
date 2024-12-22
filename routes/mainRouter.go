package routes

import (
	"github.com/gin-gonic/gin"
)

func MainRouter(api *gin.RouterGroup) {
	// since AuthRouter is under same package "routes" so no need to specifically mention we can directly write its name
	AuthRouter(api)
}
