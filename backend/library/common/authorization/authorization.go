package authorization

import (
	"backend/config/environment"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(data ClaimData) (string, error) {
	// create new claims token
	claim := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": data.JTI,
		"iss": data.ISS,
		"aud": data.AUD,
		"exp": data.EXP,
		"iat": data.IAT,
	})

	// token
	return claim.SignedString(environment.JWT_KEY)
}

func RevokeToken() {

}

func ValidateToken() {

}
