package main

import (
	"backend/config/environment"
	"backend/config/path"
	"backend/config/variable"
	"backend/library/cli"
	"backend/server"
	"fmt"
	"os"
	"slices"
)

// main applications
func main() {

	// set pathing
	path.Set()

	// load env
	environment.Save()

	// check request type
	if len(os.Args) > 1 {

		// clearly CLI because there is arguments
		variable.RequestType = "CLI"

		// check if contain allowed arguments
		if slices.Contains(variable.CLIAllowedArguments, os.Args[1]) {
			// send to cli
			cli.Run(os.Args[1])
		} else {
			// wrong arguments, send error to cli
			fmt.Println(fmt.Errorf("wrong CLI arguments, use only supported argument: %v", variable.CLIAllowedArguments))
		}

		return
	}

	// set to app
	variable.RequestType = "APP"

	// run server
	server.Start()
}
