package roleController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	// process
	defaultModel := &model.Role{}
	defaultTableName := "role"
	defaultIDColumn := "id"
	defaultSelect := []string{"id", "name", "is_superadmin", "created_at", "updated_at"}
	defaultOrderColumn := clause.OrderByColumn{
		Column: clause.Column{
			Table: defaultTableName,
			Name:  "name",
		},
	}

	// aliaeses
	defaultAliases := map[string]string{}

	// result data
	var totalUnfiltered int64
	var totalFiltered int64
	var rows []model.Module
	var result []model.Module

	// find total and filtered total
	database.CONN.Model(defaultModel).Count(&totalUnfiltered)

	// check if input query is supplied
	if input.Query != nil {
		// orm
		ormTotal := database.CONN.Model(defaultModel)

		if len(input.ExcludeID) > 0 {
			ormTotal = ormTotal.Where(fmt.Sprintf("%s NOT IN ?", defaultIDColumn), input.ExcludeID)
		}

		common.ProcessIndexQuery(ormTotal, *input.Query, defaultTableName, defaultAliases)

		// get filtered
		ormTotal.Count(&totalFiltered)
	} else {
		// same as total unfiltered
		totalFiltered = totalUnfiltered
	}

	// get data
	ormSelect := database.CONN.
		Model(defaultModel).
		Select(defaultSelect)

	// set exclusion
	if len(input.ExcludeID) > 0 {
		ormSelect = ormSelect.Where(fmt.Sprintf("%s NOT IN ?", defaultIDColumn), input.ExcludeID)
	}

	// if
	if input.Query != nil {
		common.ProcessIndexQuery(ormSelect, *input.Query, defaultTableName, defaultAliases)
	}

	// set order
	orderColumnDir := (input.Order.Dir == "desc")
	isRaw := strings.Contains(input.Order.Column, ".")
	newOrderColumn := clause.OrderByColumn{
		Column: clause.Column{
			Table: defaultTableName,
			Name:  input.Order.Column,
			Raw:   isRaw,
		},
		Desc: orderColumnDir,
	}

	var orderColumn []clause.OrderByColumn
	orderColumn = append(orderColumn, newOrderColumn, defaultOrderColumn)

	// get data
	dbSelect := ormSelect.
		Order(orderColumn).
		Limit(int(input.Limit)).
		Offset(int(input.Offset)).
		Find(&rows)

	if dbSelect.Error != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menarik data",
			"errors": gin.H{
				"query": []string{"Ada kesalahan pada query"},
			},
		})
		return
	}

	// mutate the data or just straight append
	result = append(result, rows...)

	// return data
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data": gin.H{
			"totalUnfiltered": totalUnfiltered,
			"totalFiltered":   totalFiltered,
			"rows":            result,
		},
	})
}

// save modules
func SaveModules(c *gin.Context) {
	// get and validate input
	var input RoleSaveModulesForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menyimpan data modul untuk role",
			"errors":  common.ConvertValidationError(err.Error(), RoleSaveModulesError),
		})
		return
	}

	// delete modules first from role, if exist
	database.CONN.Where("role_id = ?", input.ID).Delete(&model.RoleModuleList{})

	// find the added modules
	type MinimalModule struct {
		Name string `gorm:"column:name"`
	}

	var modules []MinimalModule
	database.CONN.
		Model(&model.Module{}).
		Select("name").
		Where("name IN ?", input.Modules).
		Find(&modules)

	// if empty then don't add anything as it means
	// the user drop all allowed modules on roles
	if len(modules) > 0 {
		// create inserted data
		var insertData []model.RoleModuleList

		// foreach on modules
		for _, module := range modules {
			insertData = append(insertData, model.RoleModuleList{
				RoleID:     input.ID,
				ModuleName: module.Name,
			})
		}

		// insert batch
		if insert := database.CONN.Create(&insertData); insert.Error != nil {
			c.AbortWithStatusJSON(503, gin.H{
				"status":  "error",
				"message": "Gagal menyimpan data modul untuk role. Database sedang sibuk.",
			})
			return
		}
	}

	// find role
	var role PartialRoleData
	database.CONN.
		Model(&model.Role{}).
		Select("name").
		Where("id = ?", input.ID).
		First(&role)

	// return ok
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Data modul untuk role %s berhasil disimpan.", role.Name),
	})
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
