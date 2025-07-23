package environment

import (
	"backend/config/variable"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// main environment variable. Typically would be between 'development' or 'production'
var ENV string

// name of the app
var APP_NAME string

// server host
var SERVER_HOST string

// server port
var SERVER_PORT string

// database type, between MySQL and Postgres
var DB_TYPE string

// database host
var DB_HOST string

// database port
var DB_PORT string

// database name
var DB_NAME string

// database auth username
var DB_USERNAME string

// database auth password
var DB_PASSWORD string

// database timezone. Typically this should be UTC
var DB_TIMEZONE string

// database engine. e.g. InnoDB etc.
var DB_ENGINE string

// database charset. e.g. utf8mb4, utf8, etc.
var DB_CHARSET string

// database collation e.g. utf8mb4_general_ci, utf8mb4_unicode_ci, etc.
var DB_COLLATE string

// database using ssl or not, between disable or enable
var DB_SSL string

// jwt key
var JWT_KEY string

// jwt access token expired in seconds
var JWT_ACCESS_EXPIRED int64

// jwt refresh token expired in seconds
var JWT_REFRESH_EXPIRED int64

// Save the DotEnv Configurations
func Save() {
	// load godotenv
	envPath := variable.BasePath + "/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		panic("DotEnv file is not found. path: " + envPath)
	}

	// save
	ENV = os.Getenv("ENV")
	APP_NAME = os.Getenv("APP_NAME")
	SERVER_HOST = os.Getenv("SERVER_HOST")
	SERVER_PORT = os.Getenv("SERVER_PORT")
	DB_TYPE = os.Getenv("DB_TYPE")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
	DB_ENGINE = os.Getenv("DB_ENGINE")
	DB_CHARSET = os.Getenv("DB_CHARSET")
	DB_COLLATE = os.Getenv("DB_COLLATE")
	DB_SSL = os.Getenv("DB_SSL")
	JWT_KEY = os.Getenv("JWT_KEY")

	// need to be converted
	accessExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_EXPIRED"))
	JWT_ACCESS_EXPIRED = int64(accessExpired)

	refreshExpired, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRED"))
	JWT_REFRESH_EXPIRED = int64(refreshExpired)
}
