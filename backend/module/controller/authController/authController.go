package authController

import (
	"backend/database"
	"backend/library/common"
	"backend/library/common/auth"
	"backend/module/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// authorization
func Authenticate(c *gin.Context) {
	// get input
	var input AuthLoginForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal",
			"errors":  common.ConvertValidationError(err.Error(), AuthenticateError),
		})
		return
	}

	// create interface
	var user common.FetchedUserData

	// get data
	database.CONN.Model(&model.User{}).
		Select("user.id", "user.name", "username", "email", "password", "role_id", "role.name as role_name", "user.created_at", "user.updated_at").
		Joins("JOIN role ON role_id = role.id").
		Where("username = ?", input.Username).
		First(&user)

	// match password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal",
			"errors": map[string][]string{
				"username": {
					"Username dan password tidak cocok",
				},
			},
		})
		return
	}

	// remove password
	user.Password = "(secret)"

	// password matched! generate token

	// create access token
	accessToken := auth.CreateAccessToken(model.User{Username: user.Username, RoleID: user.RoleID})
	if accessToken.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal, Ada kesalahan pada server",
		})
		return
	}

	// set refresh token
	refreshToken := auth.CreateRefreshToken(
		model.User{Username: user.Username, RoleID: user.RoleID, ID: user.ID},
		accessToken.Data.JTI,
	)

	if refreshToken.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal, Ada kesalahan pada server",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Otorisasi berhasil!",
		"data": gin.H{
			"accessToken":  accessToken.Token,
			"refreshToken": refreshToken.Token,
			"user":         user,
		},
	})
}

// check
func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Anda sudah terotentikasi",
	})
}

// Refresh Token Endpoint
func RefreshToken(c *gin.Context) {
	// get input and validate
	var input AuthRefreshTokenForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Rotasi token gagal",
			"errors":  common.ConvertValidationError(err.Error(), RefreshTokenError),
		})
		return
	}

	// validasi refresh token
	refreshTokenValid, err := auth.ValidateRefreshToken(input.RefreshToken, input.AccessToken)
	if !refreshTokenValid || err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Rotasi token gagal. Salah satu dari access token atau refresh token tidak valid.",
		})
		return
	}

	// type
	type userData struct {
		ID       uint64 `gorm:"column:id"`
		Username string `gorm:"column:username"`
		RoleID   uint64 `gorm:"column:role_id"`
	}

	// get user id
	var user userData
	database.CONN.
		Table("user USE INDEX(username)").
		Model(&model.User{}).
		Select("id", "role_id").
		Where("username = ?", auth.JWT_DATA.SUB).
		First(&user)

	// create new token
	newJWT := auth.RefreshToken(
		auth.JWT_DATA.JTI,
		model.User{Username: auth.JWT_DATA.SUB, RoleID: user.RoleID, ID: user.ID},
	)

	// if success
	if newJWT.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Rotasi token gagal. Ada kesalahan pada server.",
		})
		return
	}

	// return with data
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Rotasi token berhasil!",
		"data": gin.H{
			"accessToken":  newJWT.AccessToken,
			"refreshToken": newJWT.RefreshToken,
		},
	})
}

// User Register Endpoint
func Register(c *gin.Context) {
	// get input
	var input AuthRegisterForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Gagal mendaftar user baru",
			"errors":  common.ConvertValidationError(err.Error(), RegisterError),
		})
		return
	}

	// create hashed password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": "Jenis tipe data password tidak tepat",
		})
		return
	}

	// if valid then create new ORM mutate password
	var user model.User
	user.Name = input.Name
	user.Username = input.Username
	user.Email = input.Email
	user.RoleID = input.RoleID
	user.Password = string(passwordHash)

	// save
	if result := database.CONN.Create(&user); result.Error != nil {
		c.AbortWithStatusJSON(503, gin.H{
			"status":  "error",
			"message": "Database sedang sibuk, silahkan coba di lain waktu",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("User %s dengan username %s berhasil dibuat", input.Name, input.Username),
	})
}
