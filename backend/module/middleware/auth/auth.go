package auth

import (
	"backend/database"
	"backend/library/common"
	"backend/library/common/auth"
	"backend/module/model"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// check if user is logged in
func IsLoggedIn(c *gin.Context) {
	// get request headers
	headers := c.Request.Header
	authHeader, authHeaderExist := headers["Authorization"]

	// if auth header not exist
	if !authHeaderExist {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal",
			"errors": map[string][]string{
				"authorization": {
					"Anda belum melakukan otentikasi atau otorisasi",
				},
			},
		})
		return
	}

	// get bearer token
	bearerToken := strings.Replace(authHeader[0], "Bearer ", "", 1)

	// validate
	valid, err := auth.ValidateAccessToken(bearerToken)
	if !valid || err != nil {
		var message string

		// if there is wrong token in error
		// then wrong token usage
		if strings.Contains(fmt.Sprintf("%v", err), "wrong token") {
			message = "Anda hanya boleh menggunakan access token sebagai Authorization"
		} else {
			message = "Token tidak valid atau sudah kedaluwarsa"
		}

		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal",
			"errors": map[string][]string{
				"authorization": {
					message,
				},
			},
		})
		return
	}

	// if valid then get user data
	var user common.FetchedUserData
	database.CONN.Model(&model.User{}).
		Select("user.id", "user.name", "username", "email", "password", "role_id", "role.name as role_name", "user.created_at", "user.updated_at").
		Joins("JOIN role ON role_id = role.id").
		Where("username = ?", auth.JWT_DATA.SUB).
		First(&user)

	// store user data in context
	c.Set("userdata", user)

	// next
	c.Next()
}

// check if user can access this route
func CheckRole(c *gin.Context) {
	// get route name
	routeName, routeNameExist := c.Get("routeName")

	// if not exist, then not named, thus it is public
	if !routeNameExist {
		c.Next()
		return
	}

	// if route name is registered in modules
	var module model.Module
	database.CONN.
		Where("name = ?", routeName).
		First(&module)

	// check, and if not exist then it is not registered
	// thus it is public
	if module.ID == 0 {
		c.Next()
		return
	}

	// get user based on supplied JWT
	var user model.User
	database.CONN.
		Where("username = ?", auth.JWT_DATA.SUB).
		First(&user)

	// user should be exist

}
