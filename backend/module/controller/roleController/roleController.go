package roleController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// all role
func All(c *gin.Context) {

}

// create role
func Create(c *gin.Context) {
	// get and validate input
	var input RoleCreateForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menambah role",
			"errors":  common.ConvertValidationError(err.Error(), RoleCreateError),
		})
		return
	}

	// parse input
	var role model.Role
	role.Name = input.Name
	role.IsSuperadmin = input.IsSuperadmin

	// create
	if create := database.CONN.Create(&role); create.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Gagal menambah role. Database sedang sibuk.",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Role %s berhasil ditambahkan", role.Name),
	})
}

// delete role
func Delete(c *gin.Context) {

}

// find specific role
func Find(c *gin.Context) {

}

// role index
func Index(c *gin.Context) {

}

// save modules
func SaveModules(c *gin.Context) {

}

// update role
func Update(c *gin.Context) {
	// get and validate input
	var input RoleUpdateForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal memperbaharui data role",
			"errors":  common.ConvertValidationError(err.Error(), RoleUpdateError),
		})
		return
	}

	// parse input
	role := map[string]any{
		"name":          input.Name,
		"is_superadmin": input.IsSuperadmin,
	}

	// update
	if update := database.CONN.Model(&model.Role{}).Where("id = ?", input.ID).Updates(role); update.Error != nil {
		fmt.Println(update.Error.Error())
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Gagal memperbaharui data role. Database sedang sibuk.",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Role %s berhasil diperbaharui", input.Name),
	})
}
