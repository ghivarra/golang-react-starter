package server

import (
	"backend/config/environment"
	"backend/config/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start() {

	// load gin engine
	var engine *gin.Engine

	// load gin engine based on environment
	if environment.ENV == "development" {
		engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	}

	// load router
	engine = router.Load(engine)

	// use recovery
	engine.Use(gin.Recovery())

	// set no proxy
	// because the trusted proxy
	// should be configured in the gateway
	engine.SetTrustedProxies(nil)

	// run
	serverDetail := fmt.Sprintf("%s:%s", environment.SERVER_HOST, environment.SERVER_PORT)
	fmt.Println("Go/Gin Framework Server is running on " + serverDetail)
	engine.Run(serverDetail)
}
