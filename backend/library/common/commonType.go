package common

import "time"

var ConvertAllowedTypes []string = []string{
	"string",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float32",
	"float64",
	"bool",
}

type ErrorMessageInterface struct {
	Field    string
	Messages map[string]string
}

type FetchedUserData struct {
	ID           uint64    `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	Username     string    `gorm:"column:username"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	RoleID       uint64    `gorm:"column:role_id"`
	RoleName     string    `gorm:"column:role_name"`
	IsSuperadmin uint8     `gorm:"column:is_superadmin"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type CompleteUserData struct {
	ID             uint64
	Name           string
	Username       string
	Email          string
	Password       string
	RoleID         uint64
	RoleName       string
	IsSuperadmin   uint8
	ModulesAllowed []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
