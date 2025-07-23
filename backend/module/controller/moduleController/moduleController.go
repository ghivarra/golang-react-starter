package moduleController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// get all modules
func All(c *gin.Context) {
	// get model
	var modules []model.Module
	database.CONN.Order("name ASC").Find(&modules)

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    modules,
	})
}

// create module
func Create(c *gin.Context) {
	// get and validate input
	var input ModuleCreateForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal membuat modul baru",
			"errors":  common.ConvertValidationError(err.Error(), ModuleCreateError),
		})
		return
	}

	// insert module
	var module model.Module
	module.Name = input.Name
	module.Alias = input.Alias

	// check
	if result := database.CONN.Create(&module); result.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Database sedang sibuk, silahkan coba di lain waktu",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Modul %s berhasil dibuat", input.Name),
	})
}

// delete module
func Delete(c *gin.Context) {
	// get and validate input
	var input ModuleSingleForm
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menghapus modul",
			"errors":  common.ConvertValidationError(err.Error(), ModuleSingleError),
		})
		return
	}

	// delete and check status
	if delete := database.CONN.Delete(&model.Module{}, "name", input.Name); delete.Error != nil {
		c.JSON(503, gin.H{
			"status":  "error",
			"message": "Database sedang sibuk",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Modul %s berhasil dihapus", input.Name),
	})
}

// find specific module
func Find(c *gin.Context) {
	// get and validate input
	var input ModuleSingleForm
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menarik data modul",
			"errors":  common.ConvertValidationError(err.Error(), ModuleSingleError),
		})
		return
	}

	// return
	var module model.Module
	database.CONN.First(&module, "name", input.Name)

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    module,
	})
}

// fetch list of modules
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
	defaultModel := &model.Module{}
	defaultTableName := "module"
	defaultIDColumn := "name"
	defaultSelect := []string{"name", "alias", "created_at", "updated_at"}
	defaultOrderColumn := clause.OrderByColumn{
		Column: clause.Column{
			Table: defaultTableName,
			Name:  "alias",
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

// update module
func Update(c *gin.Context) {
	// get and validate input
	var input ModuleUpdateForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal memperbaharui modul " + input.Name,
			"errors":  common.ConvertValidationError(err.Error(), ModuleUpdateError),
		})
		return
	}

	// update alias
	update := database.CONN.
		Model(&model.Module{}).
		Where("name = ?", input.Name).
		Update("alias", input.Alias)

	if update.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Database sedang sibuk. Gagal memperbaharui modul " + input.Name,
		})
		return
	}

	// return ok
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Modul %s berhasil diperbaharui", input.Name),
	})
}
