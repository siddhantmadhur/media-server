package config

import (
	"context"

	"github.com/siddhantmadhur/media-server/storage"

	_ "github.com/mattn/go-sqlite3"
)

func finishedSetup() (bool, error) {

	db, queries, err := storage.GetConn()
	defer db.Close()
	if err != nil {
		return false, err
	}
	count, err := queries.IsFinishedSetup(context.Background())
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
