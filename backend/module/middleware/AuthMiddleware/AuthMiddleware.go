package authMiddleware

import (
	"backend/library/common/auth"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
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

	// next
	c.Next()
}
