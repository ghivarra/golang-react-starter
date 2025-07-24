package api

import (
	"backend/module/controller/accountController"
	"backend/module/controller/authController"
	"backend/module/controller/menuController"
	"backend/module/controller/moduleController"
	"backend/module/controller/roleController"
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
	api.OPTIONS("/*any", cors.Run)

	// auth group
	apiAuth := api.Group("auth")
	apiAuth.POST("/register", authController.Register)
	apiAuth.POST("/authenticate", authController.Authenticate)
	apiAuth.POST("/refresh-token", authController.RefreshToken)
	apiAuth.GET("/check", auth.IsLoggedIn, authController.Check)

	// account
	apiAccount := api.Group("account")
	apiAccount.Use(auth.IsLoggedIn)
	apiAccount.GET("/find", name.Save("account.find"), auth.CheckRole, accountController.Find)
	apiAccount.POST("/index", name.Save("account.index"), auth.CheckRole, accountController.Index)
	apiAccount.POST("/create", name.Save("account.create"), auth.CheckRole, authController.Register)
	apiAccount.PATCH("/change-password", name.Save("account.change-password"), auth.CheckRole, accountController.ChangePassword)
	apiAccount.PATCH("/update", name.Save("account.update"), auth.CheckRole, accountController.Update)
	apiAccount.DELETE("/activation-status", name.Save("account.activation-status"), auth.CheckRole, accountController.ActivationStatus)
	apiAccount.DELETE("/purge", name.Save("account.purge"), auth.CheckRole, accountController.Purge)

	// self user endpoint
	apiUser := api.Group("user")
	apiUser.Use(auth.IsLoggedIn)
	apiUser.GET("/", userController.Get)
	apiUser.PATCH("/change-password", userController.ChangePassword)
	apiUser.PATCH("/update", userController.Update)
	apiUser.DELETE("/deactivate", userController.Deactivate)

	// module group
	apiModule := api.Group("module")
	apiModule.Use(auth.IsLoggedIn)
	apiModule.GET("/", name.Save("module.all"), auth.CheckRole, moduleController.All)
	apiModule.GET("/find", name.Save("module.find"), auth.CheckRole, moduleController.Find)
	apiModule.POST("/index", name.Save("module.index"), auth.CheckRole, moduleController.Index)
	apiModule.POST("/create", name.Save("module.create"), auth.CheckRole, moduleController.Create)
	apiModule.PATCH("/update", name.Save("module.update"), auth.CheckRole, moduleController.Update)
	apiModule.DELETE("/delete", name.Save("module.delete"), auth.CheckRole, moduleController.Delete)

	// menu group
	apiMenu := api.Group("menu")
	apiMenu.Use(auth.IsLoggedIn)
	apiMenu.GET("/", name.Save("menu.all"), auth.CheckRole, menuController.All)
	apiMenu.GET("/find", name.Save("menu.find"), auth.CheckRole, menuController.Find)
	apiMenu.POST("/create", name.Save("menu.create"), auth.CheckRole, menuController.Create)
	apiMenu.PATCH("/update", name.Save("menu.update"), auth.CheckRole, menuController.Update)
	apiMenu.DELETE("/delete", name.Save("menu.delete"), auth.CheckRole, menuController.Delete)

	// module group
	apiRole := api.Group("role")
	apiRole.Use(auth.IsLoggedIn)
	apiRole.GET("/", name.Save("role.all"), auth.CheckRole, roleController.All)
	apiRole.GET("/find", name.Save("role.find"), auth.CheckRole, roleController.Find)
	apiRole.POST("/index", name.Save("role.index"), auth.CheckRole, roleController.Index)
	apiRole.POST("/create", name.Save("role.create"), auth.CheckRole, roleController.Create)
	apiRole.PATCH("/update", name.Save("role.update"), auth.CheckRole, roleController.Update)
	apiRole.PUT("/save-modules", name.Save("role.save-modules"), auth.CheckRole, roleController.SaveModules)
	apiRole.DELETE("/delete", name.Save("role.delete"), auth.CheckRole, roleController.Delete)

	// return instance
	return router
}
