package accountController

// partial for response message only
type AccountDataPartial struct {
	ID   uint64 `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
