package menuController

import (
	"backend/database"
	"backend/module/model"

	"github.com/gin-gonic/gin"
)

// get all menus
func All(c *gin.Context) {
	// get menus sort by sort_number
	var menu []model.Menu
	database.CONN.
		Table("menu USE INDEX (menu_sort_number)").
		Order("sort_number asc").
		Find(&menu)

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    menu,
	})
}

// create menu
func Create(c *gin.Context) {

}

// delete menu
func Delete(c *gin.Context) {

}

// find specific menu
func Find(c *gin.Context) {

}

// update menu
func Update(c *gin.Context) {

}
