package menuController

import (
	"backend/database"
	"backend/library/common"
	"backend/module/model"
	"fmt"

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
	// get input and validate
	var input MenuCreateForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal membuat menu",
			"errors":  common.ConvertValidationError(err.Error(), MenuCreateError),
		})
		return
	}

	// set input data
	var menu model.Menu
	menu.Alias = input.Alias
	menu.RouteName = input.RouteName
	menu.SortNumber = input.SortNumber

	// if not nil then insert the icon
	// if not we use the default which is null
	if input.Icon != nil {
		menu.Icon = input.Icon
	}

	// create menu
	if create := database.CONN.Create(&menu); create.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Gagal membuat menu, database sedang sibuk",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Menu %s berhasil dibuat di urutan ke-%d", menu.Alias, menu.SortNumber),
	})
}

// delete menu
func Delete(c *gin.Context) {
	// get input and validate
	var input MenuSingleForm
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menghapus menu",
			"errors":  common.ConvertValidationError(err.Error(), MenuSingleError),
		})
		return
	}

	// get
	var menu PartialMenuData
	database.CONN.
		Model(&model.Menu{}).
		Select("alias").
		Where("id = ?", input.ID).
		First(&menu)

	//delete
	if delete := database.CONN.Delete(&model.Menu{}, input.ID); delete.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Gagal menghapus menu, database sedang sibuk",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Menu %s berhasil dihapus", menu.Alias),
	})
}

// find specific menu
func Find(c *gin.Context) {
	// get input and validate
	var input MenuSingleForm
	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal menarik data menu",
			"errors":  common.ConvertValidationError(err.Error(), MenuSingleError),
		})
		return
	}

	// get data
	var menu model.Menu
	database.CONN.First(&menu, input.ID)

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    menu,
	})
}

// update menu
func Update(c *gin.Context) {
	// get input and validate
	var input MenuUpdateForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal memperbaharui menu",
			"errors":  common.ConvertValidationError(err.Error(), MenuUpdateError),
		})
		return
	}

	// set data
	updateData := map[string]any{
		"alias":      input.Alias,
		"route_name": input.RouteName,
		"icon":       input.Icon,
	}

	// update
	if update := database.CONN.Model(&model.Menu{}).Where("id = ?", input.ID).Updates(updateData); update.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Gagal menghapus menu, database sedang sibuk",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Menu %s berhasil diperbaharui", input.Alias),
	})
}
