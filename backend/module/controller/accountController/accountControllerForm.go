package accountController

// account update form
type AccountUpdateForm struct {
	ID       uint64 `json:"id" binding:"required,is_not_unique=user:id"`
	Name     string `json:"name" binding:"required,max=100"`
	Username string `json:"username" binding:"required,max=100,alphanumeric_dash,is_unique=user:username:id:ID:uint64"`
	Email    string `json:"email" binding:"required,max=100,email,is_unique=user:email:id:ID:uint64"`
	RoleID   uint64 `json:"role_id" binding:"required,is_not_unique=role:id"`
}

// query activation status
type AccountActivationQuery struct {
	ID     uint64 `form:"id" binding:"required,is_not_unique=user:id"`
	Status string `form:"status" binding:"required,in_list=activate:deactivate"`
}
