package userController

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  "error",
		"message": "Anda belum terotorisasi",
	})
}
