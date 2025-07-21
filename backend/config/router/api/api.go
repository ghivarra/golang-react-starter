package api

import (
	"backend/module/controller/authController"
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
	api.Use(cors.Run, database.Connect)

	// auth group
	apiAuth := api.Group("auth")
	apiAuth.POST("/register", authController.Register)
	apiAuth.POST("/authenticate", authController.Authenticate)
	apiAuth.POST("/refresh-token", authController.RefreshToken)

	// account
	apiAccount := api.Group("account").Use(auth.IsLoggedIn)
	apiAccount.PATCH("/update", name.Save("user.update"), auth.CheckRole, userController.Update)
	apiAccount.PATCH("/change-password", name.Save("user.change-password"), auth.CheckRole, userController.Update)
	apiAccount.DELETE("/delete", name.Save("user.delete"), auth.CheckRole, userController.Delete)

	// self user endpoint
	apiUser := api.Group("user").Use(auth.IsLoggedIn)
	apiUser.GET("/self", userController.Self)
	apiUser.PATCH("/self/change-password", userController.Self)
	apiUser.PATCH("/self/update", userController.Self)
	apiUser.POST("/self/delete", userController.Self)

	// module group
	apiModule := api.Group("module")
	apiUser.Use(auth.IsLoggedIn)
	apiModule.GET("/", moduleController.Get)
	apiModule.GET("/find", moduleController.Find)
	apiModule.POST("/create", moduleController.Create)
	apiModule.DELETE("/delete", moduleController.Delete)
	apiModule.PATCH("/update", moduleController.Update)

	// return instance
	return router
}
