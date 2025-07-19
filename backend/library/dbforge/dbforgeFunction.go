package dbforge

import (
	"fmt"
	"strings"
)

// build query for creating column inside CREATE TABLE (...) statement
func buildColumnQuery(column TableColumn) string {
	var strColumnType string
	if column.Length != nil && *column.Length > 0 {
		strColumnType = fmt.Sprintf("%s(%d)", strings.ToUpper(column.Type), column.Length)
	} else {
		strColumnType = strings.ToUpper(column.Type)
	}

	if column.IsUnsigned != nil && *column.IsUnsigned {
		strColumnType += " UNSIGNED"
	}

	var columnContext []string

	if column.IsUnique != nil && *column.IsUnique {
		columnContext = append(columnContext, "UNIQUE KEY")
	}

	if column.IsPrimaryIndex != nil && *column.IsPrimaryIndex {
		columnContext = append(columnContext, "PRIMARY KEY")
	}

	if column.IsAutoIncrement != nil && *column.IsAutoIncrement {
		columnContext = append(columnContext, "AUTO_INCREMENT")
	}

	if column.IsNullable == nil || (column.IsNullable != nil && !*column.IsNullable) {
		columnContext = append(columnContext, "NOT NULL")
	}

	if column.Default != nil && len(*column.Default) > 0 {
		columnContext = append(columnContext, fmt.Sprintf("DEFAULT %s", *column.Default))
	}

	// join all
	query := fmt.Sprintf("`%s` %s %s", column.Name, strColumnType, strings.Join(columnContext, " "))

	return query
}

// build query for creating indexes inside CREATE TABLE (...) statement
func buildColumnIndex(index TableIndex, tableName string) string {
	return fmt.Sprintf("KEY.`%s_%s` (`%s`)", tableName, index.Name, index.Name)
}

// build query for creating foreign keys inside CREATE TABLE (...) statement
func buildForeignKey(fk TableForeignKey) string {
	query := fmt.Sprintf("CONSTRAINT %s FOREIGN KEY (`%s`) REFERENCES %s(%s)", fk.Name, fk.Column, fk.RefTable, fk.RefColumn)

	if fk.OnUpdate != nil && len(*fk.OnUpdate) > 0 {
		query += " ON UPDATE " + strings.ToUpper(*fk.OnUpdate)
	}

	if fk.OnDelete != nil && len(*fk.OnDelete) > 0 {
		query += " ON DELETE " + strings.ToUpper(*fk.OnDelete)
	}

	return query
}

// build CREATE TABLE (...) query
func buildQuery(table Table, option Option) string {
	var query string
	var columns []string
	var indexes []string
	var fks []string

	// create opening query
	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (", table.Name)

	// create columns query and join
	for _, column := range table.Columns {
		columns = append(columns, buildColumnQuery(column))
	}
	query += strings.Join(columns, ", ")

	// create indexes and join
	if len(table.Indexes) > 0 {
		query += ", "
		for _, index := range table.Indexes {
			indexes = append(indexes, buildColumnIndex(index, table.Name))
		}
		query += strings.Join(indexes, ", ")
	}

	// create foreign keys / fks
	if len(table.ForeignKeys) > 0 {
		query += ", "
		for _, fk := range table.ForeignKeys {
			fks = append(fks, buildForeignKey(fk))
		}
		query += strings.Join(fks, ", ")
	}

	// closing
	query += fmt.Sprintf(") ENGINE=%s CHARSET=%s COLLATE=%s", option.Engine, option.Charset, option.Collate)

	// return
	return query
}
