package userController

import "github.com/gin-gonic/gin"

type RegisterForm struct {
	Name string `json:"name" binding:"required,max:100"`
}

func Register(c *gin.Context) {

}

func Index(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  "error",
		"message": "Anda belum terotorisasi",
	})
}
