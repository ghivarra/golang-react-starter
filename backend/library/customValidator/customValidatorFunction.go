package customValidator

import (
	"backend/database"
	"fmt"
)

func checkUnique(table string, column string, value any) (bool, error) {
	// initiate result pointer
	var result UniqueResult

	// build query
	query := fmt.Sprintf("SELECT COUNT(*) as `total` from `%s` WHERE `%s` = ?", table, column)

	// fetch data
	db := database.CONN.Raw(query, value).Scan(&result)

	// check error
	if db.Error != nil {
		return false, db.Error
	}

	// check
	isUnique := result.total < 1

	// check
	return isUnique, nil
}

func checkUniqueExcept(table string, column string, value any, columnException string, valueException any) (bool, error) {
	// initiate result pointer
	var result UniqueResult

	// build query
	query := fmt.Sprintf("SELECT COUNT(*) as `total` from `%s` WHERE `%s` = ? AND `%s` <> ?", table, column, columnException)

	// fetch data
	db := database.CONN.Raw(query, value, valueException).Scan(&result)

	// check error
	if db.Error != nil {
		return false, db.Error
	}

	// check
	isUnique := result.total < 1

	// check
	return isUnique, nil
}
