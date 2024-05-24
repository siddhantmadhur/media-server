package library

import (
	"io/fs"
	"path/filepath"
)

func ScanLibrary(libraryId int64) error {

	filepath.Walk("/", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			_, err := FFprobe(path)
			if err != nil {
				return err
			}
		}
		return err
	})

	return nil
}
