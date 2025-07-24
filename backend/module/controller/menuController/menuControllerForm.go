package menuController

type MenuCreateForm struct {
	Alias      string  `json:"alias" binding:"required,max=200"`
	RouteName  string  `json:"route_name" binding:"required,max=200"`
	SortNumber int     `json:"sort_number" binding:"required,numeric"`
	Icon       *string `json:"icon" binding:"omitnil,max=100"`
}

type MenuSingleForm struct {
	ID uint64 `json:"id" form:"id" binding:"required,is_not_unique=menu:id"`
}

type MenuUpdateForm struct {
	ID        uint64  `json:"id" binding:"required,is_not_unique=menu:id"`
	Alias     string  `json:"alias" binding:"required,max=200"`
	RouteName string  `json:"route_name" binding:"required,max=200"`
	Icon      *string `json:"icon" binding:"omitnil,max=100"`
}
