package tables

import (
	"backend/library/migration/tables/menuTable"
	"backend/library/migration/tables/moduleTable"
	"backend/library/migration/tables/roleMenuListTable"
	"backend/library/migration/tables/roleModuleListTable"
	"backend/library/migration/tables/roleTable"
	"backend/library/migration/tables/tokenRefreshTable"
	"backend/library/migration/tables/tokenRevokedTable"
	"backend/library/migration/tables/userTable"
)

// the table Up function here will be migrated
func MigrationUp() {
	moduleTable.Up()
	roleTable.Up()
	roleModuleListTable.Up()
	userTable.Up()
	tokenRefreshTable.Up()
	tokenRevokedTable.Up()
	menuTable.Up()
	roleMenuListTable.Up()
}

// the table Down function here will be reverted
func MigrationDown() {
	roleMenuListTable.Down()
	menuTable.Down()
	tokenRevokedTable.Down()
	tokenRefreshTable.Down()
	userTable.Down()
	roleModuleListTable.Down()
	roleTable.Down()
	moduleTable.Down()
}
