package main

import (
	"backend/config/environment"
	"backend/config/path"
	"backend/server"
)

func main() {

	// set pathing
	path.Set()

	// load env
	environment.Save()

	// run server
	server.Start()
}
