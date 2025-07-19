package migration

import (
	"backend/database"
	"backend/library/common/pointer"
	"backend/library/dbforge"
	"fmt"
	"os"
)

// DB Checking Table query result from COUNT(*) as total
type checkResult struct {
	Total int
}

// Checking if the table is already exist
func migrationTableExist() bool {
	dbName := os.Getenv("DB_NAME")
	rawQuery := fmt.Sprintf("SELECT COUNT(*) AS total FROM information_schema.tables WHERE table_schema = '%s' AND table_name = 'migration'", dbName)

	var result checkResult
	database.CONN.Raw(rawQuery).Scan(&result)

	return result.Total > 1
}

// create migration table only if the table is not exist
func createMigrationTable() {

	// only create if not exist
	if !migrationTableExist() {

		// create bool pointer for optional bbolean
		isUsed := true

		// migration table option
		migrationTable := dbforge.Table{
			Name: "migration",
			Columns: []dbforge.TableColumn{
				{Name: "id", Type: "bigint", IsUnsigned: &isUsed, IsPrimaryIndex: &isUsed, IsAutoIncrement: &isUsed},
				{Name: "table", Type: "varchar", Length: pointer.IntPtr(200), IsUnique: &isUsed},
				{Name: "status", Type: "varchar", Length: pointer.IntPtr(4)},
				{Name: "created_at", Type: "datetime", Default: pointer.StringPtr("CURRENT_TIMESTAMP")},
				{Name: "updated_at", Type: "datetime", Default: pointer.StringPtr("CURRENT_TIMESTAMP")},
			},
		}

		// migrate
		dbforge.CreateTable(migrationTable)
	}
}
