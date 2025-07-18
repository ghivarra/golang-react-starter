package corsMiddleware

import "github.com/gin-gonic/gin"

// run cors middleware
func Run(c *gin.Context) {

	// set header for cors
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")

	// if options then send 204 - no content
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	// next
	c.Next()
}
