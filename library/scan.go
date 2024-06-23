package library

import (
	"context"
	"database/sql"
	"errors"
	"io/fs"
	"ocelot/config"
	"ocelot/storage"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ScanLibrary(mediaId int64, cfg *config.Config) error {
	conn, queries, err := storage.GetConn(cfg)
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
			_, err := FFprobe(path)

			if err != nil {
				return nil
			}

			tokens := strings.Split(d.Name(), ".")

			if library.MediaType == "series" {
				showName, season, epsiode, err := getShowInformation(strings.ReplaceAll(path, library.DevicePath, ""), d.Name())
				if err != nil {
					return nil
				}
				err = queries.AddNewContentFile(context.Background(), storage.AddNewContentFileParams{
					MediaLibraryID: mediaId,
					CreatedAt:      time.Now(),
					FilePath:       path,
					Extension:      tokens[len(tokens)-1],
					Name:           strings.Join(tokens[:len(tokens)-1], "."),
					Title:          showName,
					SeasonNo:       sql.NullInt64{Int64: int64(season), Valid: true},
					EpisodeNo:      sql.NullInt64{Int64: int64(epsiode), Valid: true},
				})
			} else if library.MediaType == "movies" {
				err = queries.AddNewContentFile(context.Background(), storage.AddNewContentFileParams{
					MediaLibraryID: mediaId,
					CreatedAt:      time.Now(),
					FilePath:       path,
					Extension:      tokens[len(tokens)-1],
					Name:           strings.Join(tokens[:len(tokens)-1], "."),
				})
			}

			return err

		}
		return err
	})

	return err
}

// Returns show name, season no episode no and error
func getShowInformation(fullPath string, name string) (string, int, int, error) {
	tokens := strings.Split(fullPath, "/")
	if len(tokens) < 3 {
		return "", 0, 0, errors.New("Not enough information")
	}
	getNumber, err := regexp.Compile("[0-9]+")
	if err != nil {
		return "", 0, 0, err
	}

	getSeasonString, err := regexp.Compile("s[0-9]+|S[0-9]+|Season [0-9]+")
	getEpisodeString, err := regexp.Compile("e[0-9]+|E[0-9]+|Episode [0-9]+")

	if err != nil {
		return "", 0, 0, err
	}

	seasonString := getSeasonString.FindString(fullPath)
	episodeString := getEpisodeString.FindString(name)

	season, err := strconv.Atoi(getNumber.FindString(seasonString))
	episode, err := strconv.Atoi(getNumber.FindString(episodeString))

	return strings.ReplaceAll(tokens[1], ".", " "), season, episode, err

}
