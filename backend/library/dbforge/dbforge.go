package dbforge

import (
	"backend/database"
	"fmt"
	"os"
)

func CreateTable(table Table) error {
	if table.Name == "" {
		return fmt.Errorf("empty migration table")
	}

	// engine config
	var option Option
	option.Engine = os.Getenv("DB_ENGINE")
	option.Charset = os.Getenv("DB_CHARSET")
	option.Collate = os.Getenv("DB_COLLATE")

	// query
	query := buildQuery(table, option)

	// begin transaction
	db := database.CONN.Begin()

	// exec
	ctx := db.Exec(query)
	if ctx.Error != nil {
		db.Rollback()
		return ctx.Error
	}

	// commit
	db.Commit()
	return nil
}

func DropForeignKey(tableName string, fkName string) error {
	// query
	query := fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, fkName)

	// begin transaction
	db := database.CONN.Begin()

	// exec
	ctx := db.Exec(query)
	if ctx.Error != nil {
		db.Rollback()
		return ctx.Error
	}

	// commit
	db.Commit()
	return nil
}

func DropTable(tableName string) error {
	// query
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)

	// begin transaction
	db := database.CONN.Begin()

	// exec
	ctx := db.Exec(query)
	if ctx.Error != nil {
		db.Rollback()
		return ctx.Error
	}

	// commit
	db.Commit()
	return nil
}
