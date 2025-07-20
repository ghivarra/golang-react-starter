package api

import (
	"backend/module/controller/moduleController"
	"backend/module/controller/userController"
	"backend/module/middleware/corsMiddleware"
	"backend/module/middleware/dbConnectMiddleware"

	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) *gin.Engine {

	// set api router
	api := router.Group("api")
	api.Use(corsMiddleware.Run)

	// module group
	apiModule := api.Group("module")
	apiModule.GET("/", moduleController.Get)
	apiModule.GET("/find", moduleController.Find)
	apiModule.POST("/create", moduleController.Create)
	apiModule.DELETE("/delete", moduleController.Delete)
	apiModule.PATCH("/update", moduleController.Update)

	// user group
	apiUser := api.Group("user")
	apiUser.GET("/", userController.Index)
	apiUser.POST("/register", dbConnectMiddleware.Run, userController.Register)

	// return instance
	return router
}
