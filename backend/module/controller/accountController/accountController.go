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

// get user by id with partial struct
func getUser(id any) AccountDataPartial {
	// get user data
	var user AccountDataPartial
	database.CONN.Unscoped().Where("id = ?", id).First(&user)

	// return
	return user
}

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

	// data
	data := getUser(input.ID)

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
			"message": "Gagal merubah password akun",
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
	user := getUser(input.ID)

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Password akun %s berhasil diperbaharui", user.Name),
	})
}

// find specific account
func Find(c *gin.Context) {
	// get and validate input
	var input SingleIDQuery
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menarik data akun",
			"errors":  common.ConvertValidationError(err.Error(), SingleIDError),
		})
		return
	}

	// if valid then get account data
	var user common.FetchedUserData
	database.CONN.Model(&model.User{}).
		Select("user.id", "user.name", "username", "email", "password", "role_id", "role.name as role_name", "user.created_at", "user.updated_at").
		Joins("JOIN role ON role_id = role.id").
		Where("user.id = ?", input.ID).
		First(&user)

	// mask password
	user.Password = "(secret)"

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    user,
	})
}

// get index of accounts
func Index(c *gin.Context) {
	// get and validate
	var input common.IndexForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menarik data",
			"errors":  common.ConvertValidationError(err.Error(), common.IndexError),
		})
		return
	}

	fmt.Println(input)
}

// purge account, hard delete
func Purge(c *gin.Context) {
	// get input
	var input SingleIDQuery
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menghapus permanen akun",
			"errors":  common.ConvertValidationError(err.Error(), SingleIDError),
		})
		return
	}

	// get data
	user := getUser(input.ID)

	// delete data
	delete := database.CONN.Unscoped().Delete(&model.User{}, input.ID)
	if delete.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Database sedang sibuk",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Akun %s berhasil dihapus secara permanen", user.Name),
	})
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
