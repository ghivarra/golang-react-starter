package userController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User Register Endpoint
func Register(c *gin.Context) {
	var input RegisterForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal mendaftar user baru",
			"errors":  common.ConvertValidationError(err.Error(), RegisterError),
		})
		return
	}

	// create hashed password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": "Jenis tipe data password tidak tepat",
		})
		return
	}

	// if valid then create new ORM mutate password
	var user model.User
	user.Name = input.Name
	user.Username = input.Username
	user.Email = input.Email
	user.RoleID = input.RoleID
	user.Password = string(passwordHash)

	// save
	if result := database.CONN.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": "Database sedang sibuk, silahkan coba di lain waktu",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("User %s dengan username %s berhasil dibuat", input.Name, input.Username),
	})
}

// Fetch User Data Endpoint
func Index(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  "error",
		"message": "Anda belum terotorisasi",
	})
}
