package database

import (
	"backend/database"

	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	// connect
	database.Connect(false)

	// next
	c.Next()
}
