package moduleController

type ModuleCreateForm struct {
	Name  string `json:"name" binding:"required,max=200,is_unique=module:name,alphanumeric_dash"`
	Alias string `json:"alias" binding:"required,max=200"`
}

type ModuleDeleteForm struct {
	ID uint64 `json:"id" binding:"required,is_not_unique=module,id"`
}

type ModuleUpdateForm struct {
	ID    uint64 `json:"id" binding:"required,is_not_unique=module,id"`
	Name  string `json:"name" binding:"required,max=200,is_unique=module:name:id:ID:uint64,alphanumeric_dash"`
	Alias string `json:"alias" binding:"required,max=200"`
}
