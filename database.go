package main

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func runSqliteInit() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./storage.db")
	defer db.Close()
	_, err = db.ExecContext(ctx, ddl)
	if err != nil {
		return err
	}
	return nil
}
