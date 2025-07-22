package accountController

// query activation status
type AccountActivationQuery struct {
	ID     uint64 `form:"id" binding:"required,is_not_unique=user:id"`
	Status string `form:"status" binding:"required,in_list=activate:deactivate"`
}

// change password form
type AccountChangePasswordForm struct {
	ID                   uint64 `form:"id" binding:"required,is_not_unique=user:id"`
	Password             string `json:"password" binding:"required,confirmed,min=10,has_number,has_symbol,has_uppercase,has_lowercase"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

// account update form
type AccountUpdateForm struct {
	ID       uint64 `json:"id" binding:"required,is_not_unique=user:id"`
	Name     string `json:"name" binding:"required,max=100"`
	Username string `json:"username" binding:"required,max=100,alphanumeric_dash,is_unique=user:username:id:ID:uint64"`
	Email    string `json:"email" binding:"required,max=100,email,is_unique=user:email:id:ID:uint64"`
	RoleID   uint64 `json:"role_id" binding:"required,is_not_unique=role:id"`
}

// get only single id from query parameter
type SingleIDQuery struct {
	ID uint64 `json:"id" form:"id" binding:"required,is_not_unique=user:id"`
}
