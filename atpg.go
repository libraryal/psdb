package atpg

import (
	"database/sql"
	"fmt"
)

func PGConnect(info DBInfo) (*sql.DB, error) {
	db, err := sql.Open("postgres", info.DBString)
	if err != nil {
		return nil, fmt.Errorf("couldnt connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("couldnt ping database: %v", err)
	}
	return db, nil
}

func InsertOneRow(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("couldnt insert row: %v", err)
	}
	return result, nil
}

func GetOneRow(db *sql.DB, query string, args ...interface{}) (*sql.Row, error) {
	row := db.QueryRow(query, args...)
	return row, nil
}
func GetAllRows(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve rows: %w", err)
	}
	return rows, nil
}

func UpdateRow(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not update row: %w", err)
	}
	return result, nil
}

func DeleteRow(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not delete row: %w", err)
	}
	return result, nil
}
