package migration

import (
	"backend/database"
	"backend/library/customError"
	"backend/library/migration/seeder"
	"backend/library/migration/tables"
	"backend/library/migration/tablesConfig"
	"fmt"
)

func migrateTables() error {
	createMigrationTable()
	tables.MigrationUp()
	tablesConfig.AfterMigrationUp()
	return nil
}

func refreshTables() error {
	var err error

	err = resetTables()
	if err != nil {
		return err
	}

	err = migrateTables()
	if err != nil {
		return err
	}

	return nil
}

func resetTables() error {
	tables.MigrationDown()
	tablesConfig.AfterMigrationDown()
	return nil
}

func Run(command string) {
	var err error
	var msg string

	// connect DB
	database.Connect(true)

	// switch command
	switch command {
	case "migrate":
		err = migrateTables()
		msg = "All registered migration tables has been migrated."
	case "reset":
		err = resetTables()
		msg = "All registered migration tables has been dropped."
	case "refresh":
		err = refreshTables()
		msg = "All registered migration tables has been dropped and migrated again."
	case "seed":
		err = seeder.Run()
		msg = "All database seeder has been run."
	}

	// check error
	if err != nil {
		customError.SendErrorLog(msg, err)
		return
	}

	// send message to console
	fmt.Println(msg)
}
