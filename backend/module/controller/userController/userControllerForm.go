package userController

// register form type
type RegisterForm struct {
	Name                 string `json:"name" binding:"required,max=100"`
	Username             string `json:"username" binding:"required,max=100,is_unique=user-username"`
	Email                string `json:"email" binding:"required,max=100,email,is_unique=user-email"`
	RoleID               uint64 `json:"role_id" binding:"required,is_not_unique=role-id"`
	Password             string `json:"password" binding:"required,min=10,confirmed,has_uppercase,has_lowercase,has_symbol,has_number"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}
