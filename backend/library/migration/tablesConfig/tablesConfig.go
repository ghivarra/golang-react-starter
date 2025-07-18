package tablesConfig

import (
	"backend/database"
	"database/sql"
)

// list of already migrated tables
var MigratedTables []string

// Migration table ORM struct
type Migration struct {
	ID        uint         `gorm:"primaryKey;autoIncrement"`
	Table     string       `gorm:"uniqueIndex"`
	Status    string       `gorm:"not null"`
	CreatedAt sql.NullTime `gorm:"<-:create;autoCreateTime"`
	UpdatedAt sql.NullTime `gorm:"autoCreateTime;autoUpdateTime"`
}

// function to run after migration run
func AfterMigrationUp() {
	// check total
	totalTables := len(MigratedTables)

	// check if migration exist
	if totalTables > 0 {

		// migrated table
		var createdTables []Migration
		var updatedTables []Migration

		// build orm
		for _, tableName := range MigratedTables {

			// find if table already exist
			var find Migration
			database.CONN.Model(&Migration{}).Select("id", "table").Where("`table` = ?", tableName).Limit(1).First(&find)

			if find.ID != 0 {

				updatedTables = append(updatedTables, Migration{
					Table: tableName,
				})

			} else {

				createdTables = append(createdTables, Migration{
					Table:  tableName,
					Status: "Up",
				})
			}
		}

		// insert batch
		if len(createdTables) > 0 {
			database.CONN.CreateInBatches(createdTables, totalTables)
		}

		// update batch
		if len(updatedTables) > 0 {
			for _, table := range updatedTables {
				database.CONN.Model(&Migration{}).Where("`table` = ?", table.Table).UpdateColumn("status", "Up")
			}
		}
	}
}

// function to run after migration down / reverted
func AfterMigrationDown() {
	// check total
	totalTables := len(MigratedTables)

	// check if migration exist
	// and update using where IN
	if totalTables > 0 {
		database.CONN.Model(&Migration{}).Where("`table` IN ?", MigratedTables).Update("status", "Down")
	}
}
