package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"go-otoklix-blog/infra/config"
)

// Connect to database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.Get("DB_NAME"))

	if err != nil {
		return db, err
	}
	if db == nil {
		panic("DB nil")
	}
	return db, nil
}
