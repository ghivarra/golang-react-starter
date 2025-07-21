package accountController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Create new account
func Create(c *gin.Context) {

}

// change account password
func ChangePassword(c *gin.Context) {

}

// delete account
func Delete(c *gin.Context) {
	// get input id
	userID := c.DefaultQuery("id", "0")

	// get user
	type partialUser struct {
		ID   uint64 `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}

	// get user data
	var user partialUser
	database.CONN.
		Model(&model.User{}).
		Select("id", "name").
		Where("id = ?", userID).
		First(&user)

	// check if exist
	if user.Name == "" {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menghapus user",
			"errors": gin.H{
				"id": []string{"User tidak ditemukan"},
			},
		})
		return
	}

	// delete
	database.CONN.Where("id = ?", userID).Delete(&model.User{})

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Akun %s berhasil dihapus", user.Name),
	})
}

// find specific account
func Find(c *gin.Context) {

}

// get index of accounts
func Index(c *gin.Context) {

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
