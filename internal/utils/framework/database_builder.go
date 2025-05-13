package framework

import (
	"database/sql"
	"strings"
)

type DatabaseField struct {
	Name  string
	Type  string
	Extra string
}

type DatabaseTable struct {
	Name   string
	Fields []DatabaseField
}

func CreateTable(db *sql.DB, table DatabaseTable) error {
	query := "CREATE TABLE IF NOT EXISTS " + table.Name + " ("
	for _, field := range table.Fields {
		query += field.Name + " " + field.Type + " " + field.Extra + ","
	}
	query = strings.TrimSuffix(query, ",") + ")"
	_, err := db.Exec(query)
	return err
}
