package router

import (
	"backend/module/controller/userController"
	"backend/module/middleware/corsMiddleware"
	"backend/module/middleware/dbConnectMiddleware"

	"github.com/gin-gonic/gin"
)

// load the router configurations into Gin Engine
func Load(router *gin.Engine) *gin.Engine {

	// set api router
	api := router.Group("api")
	api.Use(corsMiddleware.Run)

	// user group
	apiUser := api.Group("user")
	apiUser.GET("/", userController.Index)
	apiUser.POST("/register", dbConnectMiddleware.Run, userController.Register)

	return router
}
