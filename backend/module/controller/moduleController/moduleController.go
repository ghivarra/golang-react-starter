package moduleController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"

	"github.com/gin-gonic/gin"
)

// create module
func Create(c *gin.Context) {
	// initiate variable
	var input ModuleCreateForm

	// load input
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
}

// delete module
func Delete(c *gin.Context) {

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
