package userController

// change password
type UserChangePasswordForm struct {
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	PasswordNew          string `json:"password_new" binding:"required,confirmed,min=10,has_number,has_symbol,has_uppercase,has_lowercase"`
}
