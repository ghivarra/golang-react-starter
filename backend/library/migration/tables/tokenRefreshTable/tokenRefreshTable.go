package tokenRefreshTable

import (
	"backend/library/common/pointer"
	"backend/library/dbforge"
	"backend/library/migration/tablesConfig"
	"fmt"
)

// your table name
var tableName = "token_refresh"
var foreignKey1 = fmt.Sprintf("fk_%s_user_id", tableName)

// function Up is your migration table configurations
func Up() {

	// create bool pointer for used boolean optional column
	isTrue := true

	// create table
	dbforge.CreateTable(dbforge.Table{
		Name: tableName,
		Columns: []dbforge.TableColumn{
			{Name: "id", Type: "char", Length: pointer.IntPtr(100), IsPrimaryIndex: &isTrue},
			{Name: "user_id", Type: "bigint", IsUnsigned: &isTrue},
			{Name: "expired_at", Type: "datetime"},
			{Name: "created_at", Type: "datetime", Default: pointer.StringPtr("CURRENT_TIMESTAMP")},
		},
		Indexes: []dbforge.TableIndex{
			{Name: "expired_at"},
			{Name: "user_id"},
		},
		ForeignKeys: []dbforge.TableForeignKey{
			{
				Name:      foreignKey1,
				Column:    "user_id",
				RefTable:  "user",
				RefColumn: "id",
				OnDelete:  pointer.StringPtr("cascade"),
				OnUpdate:  pointer.StringPtr("cascade"),
			},
		},
	})

	// don't remove the line below, this is to inform the migration that this table has been migrated/created
	tablesConfig.MigratedTables = append(tablesConfig.MigratedTables, tableName)
}

// function Down is in what sequence your migrated table will be removed/reverted
func Down() {
	// uncomment and delete foreign key here if exist
	dbforge.DropForeignKey(tableName, foreignKey1)

	// delete table
	dbforge.DropTable(tableName)

	// don't remove the line below, this is to inform the migration that this table has been removed
	tablesConfig.MigratedTables = append(tablesConfig.MigratedTables, tableName)
}
