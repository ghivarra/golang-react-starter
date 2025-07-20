package userController

import (
	"backend/config/environment"
	"backend/database"
	"backend/library/common"
	"backend/library/common/authorization"
	"backend/module/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// authorization
func Authenticate(c *gin.Context) {
	// get input
	var input UserLoginForm
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal",
			"errors":  common.ConvertValidationError(err.Error(), AuthorizeError),
		})
		return
	}

	// new struct
	type userData struct {
		ID        uint64    `gorm:"column:id"`
		Name      string    `gorm:"column:name"`
		Username  string    `gorm:"column:username"`
		Email     string    `gorm:"column:email"`
		Password  string    `gorm:"column:password"`
		RoleID    uint64    `gorm:"column:role_id"`
		RoleName  string    `gorm:"column:role_name"`
		CreatedAt time.Time `gorm:"column:created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
	}

	// create interface
	var user userData

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

	// get claim data
	claimData := authorization.ClaimData{
		JTI: uuid.New().String(),
		SUB: user.Username,
		AUD: user.RoleID,
		ISS: environment.APP_NAME,
		EXP: time.Now().Add(time.Second + time.Duration(environment.JWT_ACCESS_EXPIRED)).Unix(),
		IAT: time.Now().Unix(),
	}

	// password matched! return token
	accessToken, err := authorization.CreateToken(claimData)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal, Ada kesalahan pada server",
		})
		return
	}

	// set refresh token
	claimData.JTI = uuid.New().String()
	claimData.EXP = time.Now().Add(time.Second + time.Duration(environment.JWT_REFRESH_EXPIRED)).Unix()
	refreshToken, err := authorization.CreateToken(claimData)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": "Otorisasi gagal, Ada kesalahan pada server",
		})
		return
	}

	// add into refresh token table
	database.CONN.Create(&model.TokenRefresh{
		ID:        claimData.JTI,
		ExpiredAt: time.Unix(claimData.EXP, claimData.EXP*1000),
	})

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Otorisasi berhasil!",
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"user":         user,
		},
	})
}

// User Register Endpoint
func Register(c *gin.Context) {
	// get input
	var input UserRegisterForm
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

// Fetch User Data Endpoint
func Index(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  "error",
		"message": "Anda belum terotorisasi",
	})
}
