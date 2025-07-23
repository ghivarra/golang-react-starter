package database

import (
	"backend/config/environment"
	"backend/library/customError"
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// database connection instance
var CONN *gorm.DB

func buildMySQLDSN(multipleStatement bool) string {
	// set config
	host := environment.DB_HOST
	port := environment.DB_PORT
	name := environment.DB_NAME
	user := environment.DB_USERNAME
	pass := environment.DB_PASSWORD
	charset := url.QueryEscape(environment.DB_CHARSET)
	timezone := url.QueryEscape(environment.DB_TIMEZONE)

	// build dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s", user, pass, host, port, name, charset, timezone)

	// if multi true
	if multipleStatement {
		dsn += "&multiStatements=true"
	}

	// return dsn
	return dsn
}

func buildPostgresDSN() string {
	// set config
	host := environment.DB_HOST
	port := environment.DB_PORT
	name := environment.DB_NAME
	user := environment.DB_USERNAME
	pass := environment.DB_PASSWORD
	ssl := environment.DB_SSL
	timezone := environment.DB_TIMEZONE

	// build dsn
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, pass, name, port, ssl, timezone)

	// return dsn
	return dsn
}

func Connect(multipleStatement bool) {

	// init error
	var err error

	// build dsn
	var dsn string
	if environment.DB_TYPE == "postgres" {
		dsn = buildPostgresDSN()
	} else {
		dsn = buildMySQLDSN(multipleStatement)
	}

	// set logger options
	var loggerOption logger.Interface
	if environment.ENV == "development" {
		loggerOption = logger.Default.LogMode(logger.Info)
	} else {
		loggerOption = logger.Default.LogMode(logger.Error)
	}

	// connect
	CONN, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: loggerOption,
	})

	// handling error connection
	if err != nil {
		customError.SendErrorLog("Failed to connect to the database.", err)
	}
}
