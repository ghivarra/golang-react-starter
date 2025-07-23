package moduleController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// get all modules
func All(c *gin.Context) {

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
	var input ModuleDeleteForm
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menghapus modul",
			"errors":  common.ConvertValidationError(err.Error(), ModuleDeleteError),
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

}

// fetch list of modules
func Index(c *gin.Context) {

}

// update module
func Update(c *gin.Context) {

}
