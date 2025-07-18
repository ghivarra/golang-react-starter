package router

import (
	"backend/module/controller/userController"
	"backend/module/middleware/corsMiddleware"

	"github.com/gin-gonic/gin"
)

// load the router configurations into Gin Engine
func Load(router *gin.Engine) *gin.Engine {

	// set api router
	api := router.Group("api")
	api.Use(corsMiddleware.Run)
	api.GET("user", userController.Index)

	return router
}
