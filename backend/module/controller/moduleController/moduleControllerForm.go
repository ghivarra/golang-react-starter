package moduleController

type ModuleCreateForm struct {
	Name  string `json:"name" binding:"required,max=200,is_unique=module:name,alphanumeric_dash"`
	Alias string `json:"alias" binding:"required,max=200"`
}

type ModuleSingleForm struct {
	Name string `form:"name" binding:"required,is_not_unique=module:name"`
}

type ModuleUpdateForm struct {
	Name  string `json:"name" binding:"required,is_not_unique=module:name"`
	Alias string `json:"alias" binding:"required,max=200"`
}
