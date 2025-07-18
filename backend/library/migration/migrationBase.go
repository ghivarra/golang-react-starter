package migration

import (
	"backend/database"
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
		// migration table option
		migrationTable := dbforge.Table{
			Name: "migration",
			Columns: []dbforge.TableColumn{
				{Name: "id", Type: "bigint", IsUnsigned: true, IsPrimaryIndex: true, IsAutoIncrement: true},
				{Name: "table", Type: "varchar", Constraint: 200, IsUnique: true},
				{Name: "status", Type: "varchar", Constraint: 4},
				{Name: "created_at", Type: "datetime", Default: "CURRENT_TIMESTAMP", IsNullable: true},
				{Name: "updated_at", Type: "datetime", Default: "CURRENT_TIMESTAMP", IsNullable: true},
			},
		}

		// migrate
		dbforge.CreateTable(migrationTable)
	}
}
