package auth

import (
	"backend/config/environment"
	"backend/database"
	"backend/module/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var JWT_DATA ClaimData
var JWT_ACCESS_SIGN string = "ACCESS"
var JWT_REFRESH_SIGN string = "REFRESH"

func CreateToken(data ClaimData) (string, error) {
	// create new claims token
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti": data.JTI,
		"sub": data.SUB,
		"iss": data.ISS,
		"aud": data.AUD,
		"exp": data.EXP,
		"iat": data.IAT,
	})

	// token
	return claim.SignedString([]byte(environment.JWT_KEY))
}

func CreateAccessToken(user model.User) GeneratedToken {
	// mutate role_id into string for claim data
	roleID := strconv.Itoa(int(user.RoleID))
	audiences := []string{roleID}

	// create claim data
	claimData := ClaimData{
		JTI: uuid.New().String(),
		SUB: user.Username,
		AUD: audiences,
		ISS: fmt.Sprintf("%s-%s", environment.APP_NAME, JWT_ACCESS_SIGN),
		EXP: time.Now().Add(time.Minute * time.Duration(environment.JWT_ACCESS_EXPIRED)).Unix(), // 5 minutes access token
		IAT: time.Now().Unix(),
	}

	// create token
	token, err := CreateToken(claimData)

	// return
	return GeneratedToken{
		Token: token,
		Error: err,
		Data:  claimData,
	}
}

func CreateRefreshToken(user model.User, accessTokenID string) GeneratedToken {
	// mutate role_id into string for claim data
	roleID := strconv.Itoa(int(user.RoleID))
	audiences := []string{roleID, accessTokenID}

	// create claim data
	claimData := ClaimData{
		JTI: uuid.New().String(),
		SUB: user.Username,
		AUD: audiences,
		ISS: fmt.Sprintf("%s-%s", environment.APP_NAME, JWT_REFRESH_SIGN),
		EXP: time.Now().Add(time.Minute * time.Duration(environment.JWT_REFRESH_EXPIRED)).Unix(), // 30 minutes refresh token
		IAT: time.Now().Unix(),
	}

	// create token
	token, err := CreateToken(claimData)

	// add into refresh token table
	result := database.CONN.Create(&model.TokenRefresh{
		ID:        claimData.JTI,
		UserID:    user.ID,
		ExpiredAt: time.Unix(claimData.EXP, 0),
	})

	if result.Error != nil {
		return GeneratedToken{
			Token: "",
			Error: result.Error,
			Data:  claimData,
		}
	}

	// return
	return GeneratedToken{
		Token: token,
		Error: err,
		Data:  claimData,
	}
}

func RefreshToken(tokenID string, user model.User) RefreshTokenData {
	// generate new access token
	accessToken := CreateAccessToken(user)

	// if error then return the error and empty token
	if accessToken.Error != nil {
		return RefreshTokenData{
			AccessToken:  "",
			RefreshToken: "",
			Error:        accessToken.Error,
		}
	}

	// if not then create refresh token and remove old refresh token / rotate token

	// new refresh token
	refreshToken := CreateRefreshToken(user, accessToken.Data.JTI)

	// if error then return the error and empty token
	if refreshToken.Error != nil {
		return RefreshTokenData{
			AccessToken:  "",
			RefreshToken: "",
			Error:        refreshToken.Error,
		}
	}

	// remove old refresh token
	result := database.CONN.Delete(&model.TokenRefresh{}, "id", tokenID)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	// return
	return RefreshTokenData{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
		Error:        nil,
	}
}

func RevokeToken(tokenString string) {

	// validate token first to store the claim
	ValidateToken(tokenString)

	// add revoked based on jti
	database.CONN.Create(&model.TokenRevoked{
		ID:        JWT_DATA.JTI,
		ExpiredAt: time.Unix(JWT_DATA.EXP, JWT_DATA.EXP*1000),
	})
}

func storeClaimToken(claims jwt.RegisteredClaims) {
	// store the same value type
	JWT_DATA.JTI = claims.ID
	JWT_DATA.SUB = claims.Subject
	JWT_DATA.ISS = claims.Issuer
	JWT_DATA.AUD = claims.Audience
	JWT_DATA.EXP = claims.ExpiresAt.Unix()
	JWT_DATA.IAT = claims.IssuedAt.Unix()
}

func ValidateToken(tokenString string) (bool, error) {
	// set claim
	var claims jwt.RegisteredClaims
	secretKey := []byte(environment.JWT_KEY)

	// parse token
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	// store claim token
	storeClaimToken(claims)

	// return true
	return true, nil
}

func ValidateAccessToken(tokenString string) (bool, error) {
	// validate token first
	valid, err := ValidateToken(tokenString)
	if err != nil || !valid {
		return false, err
	}

	// check if access token
	if !strings.Contains(JWT_DATA.ISS, JWT_ACCESS_SIGN) {
		return false, fmt.Errorf("wrong token. Only use access token")
	}

	// return ok
	return true, nil
}

func ValidateRefreshToken(tokenString string, accessToken string) (bool, error) {
	// validate token first
	valid, err := ValidateToken(tokenString)
	if err != nil || !valid {
		return false, err
	}

	// check if access token
	if !strings.Contains(JWT_DATA.ISS, JWT_REFRESH_SIGN) {
		return false, fmt.Errorf("wrong token. Only use refresh token")
	}

	// check if in database
	var count int64
	database.CONN.
		Model(&model.TokenRefresh{}).
		Joins("INNER JOIN user ON user_id = user.id AND user.username = ?", JWT_DATA.SUB).
		Where("token_refresh.id = ?", JWT_DATA.JTI).
		Count(&count)

	// if exist
	if count < 1 {
		return false, fmt.Errorf("refresh token sudah tidak berlaku lagi")
	}

	// read access token payload without verifying because this token
	// is clearly expired
	var jwtParser jwt.Parser
	var claims jwt.RegisteredClaims
	token, _, err := jwtParser.ParseUnverified(accessToken, &claims)
	if err != nil {
		fmt.Println(token) // this is redundant but I should use the token variable
		return false, err
	}

	// validasi if the supplied access token same with payload
	accessTokenID := JWT_DATA.AUD[1]
	if accessTokenID != claims.ID {
		return false, fmt.Errorf("refresh token yang divalidasi tidak cocok dengan access token yang diberikan")
	}

	// return ok
	return true, nil
}
