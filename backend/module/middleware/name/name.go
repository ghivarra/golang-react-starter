package name

import "github.com/gin-gonic/gin"

func Save(routeName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// store name
		c.Set("routeName", routeName)
		c.Next()
	}
}
