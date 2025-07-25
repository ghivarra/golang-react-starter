package roleModuleListTable

import (
	"backend/library/common/pointer"
	"backend/library/dbforge"
	"backend/library/migration/tablesConfig"
	"fmt"
)

// your table name
var tableName = "role_module_list"
var foreignKey1 = fmt.Sprintf("fk_%s_role_id", tableName)
var foreignKey2 = fmt.Sprintf("fk_%s_module_name", tableName)

// function Up is your migration table configurations
func Up() {

	// create bool pointer for used boolean optional column
	isTrue := true

	// create table
	dbforge.CreateTable(dbforge.Table{
		Name: tableName,
		Columns: []dbforge.TableColumn{
			{Name: "id", Type: "bigint", IsUnsigned: &isTrue, IsPrimaryIndex: &isTrue, IsAutoIncrement: &isTrue},
			{Name: "role_id", Type: "bigint", IsUnsigned: &isTrue},
			{Name: "module_name", Type: "varchar", Length: pointer.IntPtr(200)},
		},
		Indexes: []dbforge.TableIndex{
			{Name: "role_id"},
			{Name: "module_name"},
		},
		ForeignKeys: []dbforge.TableForeignKey{
			{
				Name:      foreignKey1,
				Column:    "role_id",
				RefTable:  "role",
				RefColumn: "id",
				OnDelete:  pointer.StringPtr("cascade"),
				OnUpdate:  pointer.StringPtr("cascade"),
			},
			{
				Name:      foreignKey2,
				Column:    "module_name",
				RefTable:  "module",
				RefColumn: "name",
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
	dbforge.DropForeignKey(tableName, foreignKey2)

	// delete table
	dbforge.DropTable(tableName)

	// don't remove the line below, this is to inform the migration that this table has been removed
	tablesConfig.MigratedTables = append(tablesConfig.MigratedTables, tableName)
}
