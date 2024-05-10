package storage

import (
	"database/sql"
)

func GetConn() (*sql.DB, *Queries, error) {

	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		return nil, nil, err
	}
	queries := New(db)

	return db, queries, err
}
