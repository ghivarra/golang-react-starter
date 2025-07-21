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

	// if user not found then kick
	if user.Email == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal",
			"errors": map[string][]string{
				"authorization": {
					"Akun anda sudah dinonaktifkan",
				},
			},
		})
		return
	}

	// store user data in context
	c.Set("userdata", user)

	// next
	c.Next()
}

// check if user can access this route
func CheckRole(c *gin.Context) {
	// get userdata
	userData, userDataExist := c.Get("userdata")
	if !userDataExist {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": "Data user tidak ditemukan",
		})
		return
	}

	// get user based on supplied JWT
	user := userData.(common.FetchedUserData)

	// get role
	type roleData struct {
		ID           uint64 `gorm:"column:id"`
		IsSuperadmin uint   `gorm:"column:is_superadmin"`
	}

	var role roleData
	database.CONN.
		Model(&model.Role{}).
		Select("id", "is_superadmin").
		First(&role, user.RoleID)

	// check if is superadmin, if value more than 0, then always pass
	if role.IsSuperadmin > 0 {
		c.Next()
		return
	}

	// get route name
	routeName, routeNameExist := c.Get("routeName")

	// if not exist, then not named, thus it is public
	if !routeNameExist {
		c.Next()
		return
	}

	// if route name is registered in modules
	var countModule int64
	database.CONN.
		Model(&model.Module{}).
		Where("name = ?", routeName).
		Count(&countModule)

	// check, and if not exist then it is not registered
	// thus it is public
	if countModule < 1 {
		c.Next()
		return
	}

	// check if module is exist in role check
	var inList int64
	database.CONN.
		Model(&model.RoleModuleList{}).
		Where("role_id = ?", user.RoleID).
		Where("module_name = ?", routeName).
		Count(&inList)

	// if not exist then it is forbidden
	if inList < 1 {
		c.AbortWithStatusJSON(403, gin.H{
			"status":  "error",
			"message": "Anda tidak diizinkan untuk mengakses halaman ini",
		})
		return
	}

	// if not superadmin then get module list
	c.Next()
}
