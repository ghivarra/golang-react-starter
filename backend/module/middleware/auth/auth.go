package auth

import (
	"backend/database"
	"backend/library/common"
	"backend/library/common/auth"
	"backend/module/model"
	"fmt"
	"slices"
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
		Table("user USE INDEX(username)").
		Select("user.id", "user.name", "username", "email", "password", "role_id", "role.name as role_name", "role.is_superadmin", "user.created_at", "user.updated_at").
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

	// set user into new user
	var completeUser common.CompleteUserData
	completeUser.ID = user.ID
	completeUser.Name = user.Name
	completeUser.Email = user.Email
	completeUser.Username = user.Username
	completeUser.Password = user.Password
	completeUser.RoleID = user.RoleID
	completeUser.RoleName = user.RoleName
	completeUser.IsSuperadmin = user.IsSuperadmin
	completeUser.CreatedAt = user.CreatedAt
	completeUser.UpdatedAt = user.UpdatedAt

	// list module names
	type ModuleNames struct {
		Name string `gorm:"column:module_name"`
	}

	// get role modules if not superadmin
	if user.IsSuperadmin < 1 {
		// get list of modules
		var modules []ModuleNames
		database.CONN.
			Table("role_module_list USE INDEX(role_module_list_role_id)").
			Model(&model.RoleModuleList{}).
			Select("module_name").
			Where("role_id = ?", user.RoleID).
			Find(&modules)

		// put
		for _, module := range modules {
			completeUser.ModulesAllowed = append(completeUser.ModulesAllowed, module.Name)
		}
	}

	// store user data in context
	c.Set("userdata", completeUser)

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
	user := userData.(common.CompleteUserData)

	// check if is superadmin, if value more than 0, then always pass
	if user.IsSuperadmin > 0 {
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

	// check if is in list
	routeNameStr := routeName.(string)
	if !slices.Contains(user.ModulesAllowed, routeNameStr) {
		c.AbortWithStatusJSON(403, gin.H{
			"status":  "error",
			"message": "Anda tidak diizinkan untuk mengakses halaman ini",
		})
		return
	}

	// next, fine!
	c.Next()
}
