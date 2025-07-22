package accountController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// active or inactive
// we use the soft delete feature for this
func ActivationStatus(c *gin.Context) {
	// get query data
	var input AccountActivationQuery
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal merubah status akun",
			"errors":  common.ConvertValidationError(err.Error(), AccountUpdateError),
		})
		return
	}

	// get data
	type accountDataPartial struct {
		ID   uint64 `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}

	// data
	var data accountDataPartial

	// fetch if exist
	database.CONN.
		Model(&model.User{}).
		Unscoped().
		Where("id = ?", input.ID).
		First(&data)

	// variables
	var action string
	var result *gorm.DB

	// action
	if input.Status == "deactivate" {
		result = database.CONN.Delete(&model.User{}, input.ID)
		action = "dinonaktifkan"
	} else {
		result = database.CONN.Model(&model.User{}).Unscoped().Where("id = ?", input.ID).Update("deleted_at", nil)
		action = "diaktifkan"
	}

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		c.JSON(503, gin.H{
			"status":  "error",
			"message": "Database sedang sibuk",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Akun %s berhasil %s", data.Name, action),
	})
}

// change account password
func ChangePassword(c *gin.Context) {
	// get and validate input
	var input AccountChangePasswordForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal merubah status akun",
			"errors":  common.ConvertValidationError(err.Error(), AccountChangePasswordError),
		})
		return
	}

	// hash new password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
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
		Where("id = ?", input.ID).
		Update("password", string(passwordHash))

	if update.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Server sedang sibuk",
		})
		return
	}

	// get user data
	type partialData struct {
		ID   uint64 `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}

	var user partialData
	database.CONN.
		Model(&model.User{}).
		Select("id", "name").
		Where("id = ?", input.ID).
		First(&user)

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Password akun %s berhasil diperbaharui", user.Name),
	})
}

// find specific account
func Find(c *gin.Context) {

}

// get index of accounts
func Index(c *gin.Context) {

}

// purge account, hard delete
func Purge(c *gin.Context) {

}

// User Update endpoint
func Update(c *gin.Context) {
	// get and validate input
	var input AccountUpdateForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal memperbaharui data akun",
			"errors":  common.ConvertValidationError(err.Error(), AccountUpdateError),
		})
		return
	}

	// type
	type userUpdatePartial struct {
		Name     string `gorm:"column:name"`
		Username string `gorm:"column:username"`
		Email    string `gorm:"column:email"`
		RoleID   uint64 `gorm:"column:role_id"`
	}

	// update user
	result := database.CONN.
		Model(&model.User{}).
		Where("id = ?", input.ID).
		Updates(userUpdatePartial{
			Name:     input.Name,
			Username: input.Username,
			Email:    input.Email,
			RoleID:   input.RoleID,
		})

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Gagal memperbaharui data akun",
		})
		return
	}

	// send ok
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Data akun '%s' berhasil diperbaharui", input.Name),
	})
}
