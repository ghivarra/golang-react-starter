package roleController

type RoleCreateForm struct {
	Name         string `json:"name" binding:"required,is_unique=role:name,max=60"`
	IsSuperadmin int    `json:"is_superadmin" binding:"in_list=0:1"`
}

type RoleSaveModulesForm struct {
	ID      uint64   `json:"id" form:"id" binding:"required,is_not_unique=role:id"`
	Modules []string `json:"modules" binding:"required"`
}

type RoleSingleID struct {
	ID uint64 `json:"id" form:"id" binding:"required,is_not_unique=role:id"`
}

type RoleUpdateForm struct {
	ID           uint64 `json:"id" binding:"required,is_not_unique=role:id"`
	Name         string `json:"name" binding:"required,is_unique=role:name:id:ID:uint64,max=60"`
	IsSuperadmin int    `json:"is_superadmin" binding:"in_list=0:1"`
}
