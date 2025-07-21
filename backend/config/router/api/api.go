package api

import (
	"backend/module/controller/accountController"
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
	apiAccount := api.Group("account")
	apiAccount.Use(auth.IsLoggedIn)

	apiAccount.GET("/", name.Save("user.index"), auth.CheckRole, accountController.Index)
	apiAccount.GET("/find", name.Save("user.find"), auth.CheckRole, accountController.Find)
	apiAccount.POST("/create", name.Save("user.create"), auth.CheckRole, accountController.Create)
	apiAccount.PATCH("/change-password", name.Save("user.change-password"), auth.CheckRole, accountController.Update)
	apiAccount.PATCH("/update", name.Save("user.update"), auth.CheckRole, accountController.Update)
	apiAccount.DELETE("/delete", name.Save("user.delete"), auth.CheckRole, accountController.Delete)

	// self user endpoint
	apiUser := api.Group("user")
	apiUser.Use(auth.IsLoggedIn)

	apiUser.GET("/", userController.Get)
	apiUser.PATCH("/change-password", userController.ChangePassword)
	apiUser.PATCH("/update", userController.Update)
	apiUser.POST("/delete", userController.Delete)

	// module group
	apiModule := api.Group("module")
	apiModule.Use(auth.IsLoggedIn)

	apiModule.GET("/", moduleController.Index)
	apiModule.GET("/find", moduleController.Find)
	apiModule.POST("/create", moduleController.Create)
	apiModule.PATCH("/update", moduleController.Update)
	apiModule.DELETE("/delete", moduleController.Delete)

	// return instance
	return router
}
