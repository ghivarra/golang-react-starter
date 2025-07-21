package userTable

import (
	"backend/library/common/pointer"
	"backend/library/dbforge"
	"backend/library/migration/tablesConfig"
	"fmt"
)

// your table name
var tableName = "user"
var foreignKey1 = fmt.Sprintf("fk_%s_user_id", tableName)

// function Up is your migration table configurations
func Up() {

	// create bool pointer for used boolean optional column
	isTrue := true

	// create table
	dbforge.CreateTable(dbforge.Table{
		Name: tableName,
		Columns: []dbforge.TableColumn{
			{Name: "id", Type: "bigint", IsUnsigned: &isTrue, IsPrimaryIndex: &isTrue, IsAutoIncrement: &isTrue},
			{Name: "name", Type: "varchar", Length: pointer.IntPtr(100)},
			{Name: "username", Type: "varchar", Length: pointer.IntPtr(100), IsUnique: &isTrue},
			{Name: "email", Type: "varchar", Length: pointer.IntPtr(100), IsUnique: &isTrue},
			{Name: "password", Type: "varchar", Length: pointer.IntPtr(200)},
			{Name: "role_id", Type: "bigint", IsUnsigned: &isTrue},
			{Name: "created_at", Type: "datetime", Default: pointer.StringPtr("CURRENT_TIMESTAMP")},
			{Name: "updated_at", Type: "datetime", Default: pointer.StringPtr("CURRENT_TIMESTAMP")},
			{Name: "deleted_at", Type: "datetime", IsNullable: &isTrue},
		},
		Indexes: []dbforge.TableIndex{
			{Name: "role_id"},
			{Name: "deleted_at"},
		},
		ForeignKeys: []dbforge.TableForeignKey{
			{
				Name:      foreignKey1,
				Column:    "role_id",
				RefTable:  "role",
				RefColumn: "id",
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
