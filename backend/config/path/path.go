package path

import (
	"backend/config/variable"
	"os"
	"path"
)

// set path variables inside backend/config/variable module
func Set() {

	// set directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// set base path
	variable.BasePath = path.Clean(dir)
	variable.LibraryPath = path.Clean(dir + "/library")
	variable.ControllerPath = path.Clean(dir + "/module/controller")
}
