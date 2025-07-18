package path

import (
	"backend/config/variable"
	"os"
	"path"
)

func Set() {

	// set directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// set base path
	variable.BasePath = path.Clean(dir)
}
