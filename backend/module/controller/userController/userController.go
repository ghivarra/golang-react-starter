package userController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// change own password account
func ChangePassword(c *gin.Context) {
	// get and validate input
	var input UserChangePasswordForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal merubah password",
			"errors":  common.ConvertValidationError(err.Error(), UserChangePasswordError),
		})
		return
	}

	// update password
	userdata, userdataExist := c.Get("userdata")
	if !userdataExist {
		c.AbortWithStatusJSON(404, gin.H{
			"status":  "error",
			"message": "User tidak ditemukan",
		})
		return
	}

	// set
	user := userdata.(common.FetchedUserData)

	// verifikasi password lama
	invalidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if invalidPassword != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Password lama yang anda input tidak tepat",
		})
		return
	}

	// hash new password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordNew), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Server sedang sibuk",
		})
		return
	}

	// hashed password
	update := database.CONN.
		Model(&model.User{}).
		Where("id = ?", user.ID).
		Update("password", string(passwordHash))

	if update.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Server sedang sibuk",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Password akun anda berhasil diperbaharui",
	})
}

// delete own account
func Delete(c *gin.Context) {

}

// Fetch User Data Endpoint
func Get(c *gin.Context) {
	// get data
	userData, userDataExist := c.Get("userdata")
	if !userDataExist {
		c.JSON(422, gin.H{
			"status":  "error",
			"message": "User tidak ditemukan",
		})
	}

	// convert type
	user, assertOk := userData.(common.FetchedUserData)
	if !assertOk {
		c.JSON(422, gin.H{
			"status":  "error",
			"message": "Data user tidak valid",
		})
	}

	// return data
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    user,
	})
}

// update own account data
func Update(c *gin.Context) {

}
