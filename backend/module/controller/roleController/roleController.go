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
	// model
	var roles []model.Role
	database.CONN.Order("name ASC").Find(&roles)

	// return data
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    roles,
	})
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
	// get and validate input
	var input RoleSingleID
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menghapus role",
			"errors":  common.ConvertValidationError(err.Error(), RoleSingleIDError),
		})
		return
	}

	// get data first
	model := &model.Role{}
	var data PartialRoleData
	database.CONN.
		Model(model).
		Select("name").
		Where("id = ?", input.ID).
		First(&data)

	// delete
	if delete := database.CONN.Delete(model, input.ID); delete.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Gagal menghapus role. Database sedang sibuk.",
		})
		return
	}

	// done
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Role %s berhasil dihapus", data.Name),
	})
}

// find specific role
func Find(c *gin.Context) {
	// get and validate input
	var input RoleSingleID
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menarik data",
			"errors":  common.ConvertValidationError(err.Error(), RoleSingleIDError),
		})
		return
	}

	// find
	var role model.Role
	database.CONN.Find(&role, input.ID)

	// find role modules
	var modules []PartialModuleData
	database.CONN.
		Model(&model.RoleModuleList{}).
		Select("module_name").
		Order("module_name ASC").
		Where("role_id", input.ID).
		Find(&modules)

	// result
	var result CompleteRoleData
	result.ID = input.ID
	result.Name = role.Name
	result.IsSuperadmin = role.IsSuperadmin
	result.CreatedAt = role.CreatedAt
	result.UpdatedAt = role.UpdatedAt

	// append
	if len(modules) > 0 {
		for _, module := range modules {
			result.ModulesAllowed = append(result.ModulesAllowed, module.Name)
		}
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    result,
	})
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
