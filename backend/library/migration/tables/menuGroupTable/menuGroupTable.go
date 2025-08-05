package menuGroupTable

import (
	"backend/library/common/pointer"
	"backend/library/dbforge"
	"backend/library/migration/tablesConfig"
)

// your table name
var tableName = "menu_group"

// function Up is your migration table configurations
func Up() {

	// create bool pointer for used boolean optional column
	isTrue := true

	// create table
	dbforge.CreateTable(dbforge.Table{
		Name: tableName,
		Columns: []dbforge.TableColumn{
			{Name: "id", Type: "bigint", IsUnsigned: &isTrue, IsPrimaryIndex: &isTrue, IsAutoIncrement: &isTrue},
			{Name: "name", Type: "varchar", Length: pointer.IntPtr(50)},
			{Name: "sort_number", Type: "int", Default: pointer.StringPtr("1")},
		},
		Indexes: []dbforge.TableIndex{
			{Name: "sort_number"},
		},
		ForeignKeys: []dbforge.TableForeignKey{},
	})

	// don't remove the line below, this is to inform the migration that this table has been migrated/created
	tablesConfig.MigratedTables = append(tablesConfig.MigratedTables, tableName)
}

// function Down is in what sequence your migrated table will be removed/reverted
func Down() {
	// uncomment and delete foreign key here if exist
	// dbforge.DropForeignKey(tableName, "foreign_key_name")

	// delete table
	dbforge.DropTable(tableName)

	// don't remove the line below, this is to inform the migration that this table has been removed
	tablesConfig.MigratedTables = append(tablesConfig.MigratedTables, tableName)
}
