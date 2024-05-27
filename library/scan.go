package library

import (
	"context"
	"io/fs"
	"ocelot/storage"
	"path/filepath"
	"strings"
	"time"
)

func ScanLibrary(mediaId int64) error {
	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}

	library, err := queries.GetMediaLibrary(context.Background(), mediaId)
	if err != nil {
		return err
	}

	err = filepath.WalkDir(library.DevicePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			conn, queries, err := storage.GetConn()
			defer conn.Close()
			if err != nil {
				return err
			}
			tokens := strings.Split(d.Name(), ".")
			_, err = queries.AddContent(context.Background(), storage.AddContentParams{
				CreatedAt:      time.Now(),
				FilePath:       path,
				MediaLibraryID: library.ID,
				Extension:      tokens[len(tokens)-1],
				Name:           strings.Join(tokens[:len(tokens)-1], "."),
			})
			if err != nil {
				return err
			}
		}
		return err
	})

	return err
}
