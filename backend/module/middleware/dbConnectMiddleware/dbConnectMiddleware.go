package dbConnectMiddleware

import (
	"backend/database"

	"github.com/gin-gonic/gin"
)

// middleware to connect to the database
func Run(c *gin.Context) {
	// connect
	database.Connect(false)

	// next
	c.Next()
}
