package common

// index form
type IndexForm struct {
	Limit     uint           `json:"limit" binding:"required,numeric"`
	Offset    uint           `json:"offset" binding:"required,numeric"`
	ExcludeID []uint         `json:"excludeID" binding:"required"`
	Columns   []IndexColumns `json:"columns" binding:"required,dive"`
	Order     IndexOrder     `json:"order" binding:"required"`
	Query     *[]IndexQuery  `json:"query" binding:"dive"`
}

// index form columns
type IndexColumns struct {
	ColumnName       string `json:"name" binding:"required"`
	ColumnSearchable *bool  `json:"searchable" binding:"boolean"`
	ColumnSortable   *bool  `json:"sortable" binding:"boolean"`
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
