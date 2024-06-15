package storage

import (
	"database/sql"
	"ocelot/config"
	"os"
)

func GetConn(cfg *config.Config) (*sql.DB, *Queries, error) {
	os.MkdirAll(cfg.PersistentDir, 0755)
	db, err := sql.Open("sqlite3", cfg.PersistentDir+"/storage.db")
	if err != nil {
		return nil, nil, err
	}
	queries := New(db)

	return db, queries, err
}
