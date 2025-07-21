package authController

// refresh token form interface
type AuthRefreshTokenForm struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// login form interface
type AuthLoginForm struct {
	Username string `json:"username" binding:"required,is_not_unique=user:username"`
	Password string `json:"password" binding:"required"`
}

// register form interface
type AuthRegisterForm struct {
	Name                 string `json:"name" binding:"required,max=100"`
	Username             string `json:"username" binding:"required,max=100,alphanumeric_dash,is_unique=user:username"`
	Email                string `json:"email" binding:"required,max=100,email,is_unique=user:email"`
	RoleID               uint64 `json:"role_id" binding:"required,is_not_unique=role:id"`
	Password             string `json:"password" binding:"required,min=10,confirmed,has_uppercase,has_lowercase,has_symbol,has_number"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}
