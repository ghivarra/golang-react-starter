package api

import (
	"backend/module/controller/moduleController"
	"backend/module/controller/userController"
	"backend/module/middleware/auth"
	"backend/module/middleware/cors"
	"backend/module/middleware/database"
	"backend/module/middleware/name"

	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) *gin.Engine {

	// set api router
	api := router.Group("api")
	api.Use(cors.Run)
	api.Use(database.Connect)

	// module group
	apiModule := api.Group("module")
	apiModule.GET("/", moduleController.Get)
	apiModule.GET("/find", moduleController.Find)
	apiModule.POST("/create", moduleController.Create)
	apiModule.DELETE("/delete", moduleController.Delete)
	apiModule.PATCH("/update", moduleController.Update)

	// user group
	apiUser := api.Group("user")
	apiUser.POST("/register", userController.Register)
	apiUser.POST("/authenticate", userController.Authenticate)
	apiUser.POST("/refresh-token", userController.RefreshToken)
	// all user modification
	apiUser.PATCH("/update", auth.IsLoggedIn, name.Save("user.update"), auth.CheckRole, userController.Update)
	apiUser.DELETE("/delete", auth.IsLoggedIn, name.Save("user.delete"), auth.CheckRole, userController.Delete)
	// self user endpoint
	apiUser.GET("/", auth.IsLoggedIn, userController.Get)
	apiUser.PATCH("/change-password", auth.IsLoggedIn, userController.Get)
	apiUser.PATCH("/self-update", auth.IsLoggedIn, userController.Get)
	apiUser.POST("/remove-account", auth.IsLoggedIn, userController.Get)

	// return instance
	return router
}
