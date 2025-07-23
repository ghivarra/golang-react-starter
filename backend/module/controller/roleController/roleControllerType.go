package roleController

import "time"

type PartialRoleData struct {
	Name string
}

type CompleteRoleData struct {
	ID             uint64
	Name           string
	IsSuperadmin   int
	ModulesAllowed []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type PartialModuleData struct {
	Name string `gorm:"column:module_name"`
}
