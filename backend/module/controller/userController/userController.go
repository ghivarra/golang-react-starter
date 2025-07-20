package userController

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Name                 string `json:"name" binding:"required,max=100"`
	Username             string `json:"username" binding:"required,max=100,is_unique=user-username"`
	Email                string `json:"email" binding:"required,max=100,email,is_unique=user-email"`
	Password             string `json:"password" binding:"required,min=10,confirmed,has_uppercase,has_lowercase,has_symbol,has_number"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal mendaftar user baru",
			"error":   err.Error(),
		})
		return
	}

	// print all
	fmt.Println(input)
}

func Index(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  "error",
		"message": "Anda belum terotorisasi",
	})
}
