package common

// index form
type IndexForm struct {
	Limit     uint          `json:"limit" binding:"required,numeric"`
	Offset    uint          `json:"offset" binding:"numeric"`
	ExcludeID []any         `json:"excludeID" binding:"required"`
	Order     IndexOrder    `json:"order" binding:"required"`
	Query     *[]IndexQuery `json:"query" binding:"dive"`
}

// index form order
type IndexOrder struct {
	Column string `binding:"required"`
	Dir    string `binding:"required,in_list=asc:desc"`
}

// index query
type IndexQuery struct {
	QueryColumn  string `json:"column" binding:"required"`
	QueryCommand string `json:"command" binding:"required,in_list=is:is_not:contain:not_contain"`
	QueryValue   string `json:"value" binding:"required"`
}
