package userController

// change password form
type UserChangePasswordForm struct {
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	PasswordNew          string `json:"password_new" binding:"required,confirmed,min=10,has_number,has_symbol,has_uppercase,has_lowercase"`
}

// update form
type UserUpdateForm struct {
	ID       uint64 `json:"id" binding:"required,is_not_unique=user:id"`
	Name     string `json:"name" binding:"required,max=100"`
	Username string `json:"username" binding:"required,max=100,alphanumeric_dash,is_unique=user:username:id:ID:uint64"`
	Email    string `json:"email" binding:"required,max=100,email,is_unique=user:email:id:ID:uint64"`
}
