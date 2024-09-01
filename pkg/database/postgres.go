// Package database is a package that contains the database connection.
package database

import (
	"database/sql"
)

// Conn creates a database connection.
func Conn() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres host=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
