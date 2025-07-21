package userController

import (
	"backend/library/common"

	"github.com/gin-gonic/gin"
)

// change own password account
func ChangePassword(c *gin.Context) {

}

// delete own account
func Delete(c *gin.Context) {

}

// Fetch User Data Endpoint
func Get(c *gin.Context) {
	// get data
	userData, userDataExist := c.Get("userdata")
	if !userDataExist {
		c.JSON(422, gin.H{
			"status":  "error",
			"message": "User tidak ditemukan",
		})
	}

	// convert type
	user, assertOk := userData.(common.FetchedUserData)
	if !assertOk {
		c.JSON(422, gin.H{
			"status":  "error",
			"message": "Data user tidak valid",
		})
	}

	// return data
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    user,
	})
}

// update own account data
func Update(c *gin.Context) {

}
