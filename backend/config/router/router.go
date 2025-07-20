package router

import (
	"backend/config/router/api"

	"github.com/gin-gonic/gin"
)

// load the router configurations into Gin Engine
func Load(router *gin.Engine) *gin.Engine {

	// load api group router
	router = api.Load(router)

	// return router instance
	return router
}
