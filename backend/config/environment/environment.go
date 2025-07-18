package environment

import (
	"backend/config/variable"
	"os"

	"github.com/joho/godotenv"
)

// main env
var ENV string

// server env
var SERVER_HOST string
var SERVER_PORT string

// save
func Save() {
	// load godotenv
	envPath := variable.BasePath + "/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		panic("DotEnv file is not found. path: " + envPath)
	}

	// save
	ENV = os.Getenv("ENV")
	SERVER_HOST = os.Getenv("SERVER_HOST")
	SERVER_PORT = os.Getenv("SERVER_PORT")
}
