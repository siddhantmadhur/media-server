package library

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"ocelot/config"
	"ocelot/content"
	"ocelot/content/tmdb"
	"ocelot/storage"
	"os"
	"strings"
	"time"
)

func ScanMediaFiles(library storage.MediaLibrary, cfg *config.Config) error {
	client, err := content.NewClient(tmdb.Client{
		ApiKey: os.Getenv("TMDB_READ_TOKEN"),
	})
	if err != nil {
		return err
	}

	conn, queries, err := storage.GetConn(cfg)
	defer conn.Close()
	if err != nil {
		return err
	}

	if library.MediaType == "series" {
		showFiles, err := os.ReadDir(library.DevicePath)
		if err != nil {
			return err
		}

		for _, showFile := range showFiles {
			if showFile.IsDir() {
				contentFile, err := queries.GetContentFromPath(context.Background(), library.DevicePath+"/"+showFile.Name())
				if err != nil {
					results, err := client.SearchShows(content.SearchParam{
						Query: strings.ReplaceAll(showFile.Name(), ".", " "),
					})
					// Could not find show
					if err != nil || len(results.Results) == 0 {
						contentFile, err = queries.AddNewContentFile(context.Background(), storage.AddNewContentFileParams{
							MediaLibraryID: library.ID,
							CreatedAt:      time.Now(),
							FilePath:       library.DevicePath + "/" + showFile.Name(),
							Name:           showFile.Name(),
							MediaTitle:     showFile.Name(),
							MediaType:      "series",
							Classifier:     "show",
						})
						if err != nil {
							log.Printf("error: %s\n", err.Error())
						}
					} else {
						bestResult := results.Results[0]
						contentFile, err = queries.AddNewContentFile(context.Background(), storage.AddNewContentFileParams{
							MediaLibraryID:     library.ID,
							CreatedAt:          time.Now(),
							FilePath:           library.DevicePath + "/" + showFile.Name(),
							Name:               showFile.Name(),
							MediaType:          "series",
							Classifier:         "show",
							MediaTitle:         bestResult.Name,
							Description:        sql.NullString{String: bestResult.Overview, Valid: true},
							ExternalProvider:   sql.NullString{String: "tmdb", Valid: true},
							ExternalProviderID: sql.NullInt64{Int64: int64(bestResult.Id), Valid: true},
						})
						if err != nil {
							log.Printf("error: %s\n", err.Error())
						}
					}
				}
				ScanForVideos(library.DevicePath+"/"+showFile.Name(), client, contentFile, library, queries)
			}
		}
	}

	return err
}

func ScanForVideos(currentPath string, client content.Client, parentLibrary storage.ContentLibrary, mediaLibrary storage.MediaLibrary, queries *storage.Queries) error {

	files, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		childPath := currentPath + "/" + file.Name()
		if file.IsDir() {
			ScanForVideos(childPath, client, parentLibrary, mediaLibrary, queries)
		} else {
			_, err := FFprobe(childPath)
			if err != nil {
				continue
			}
			_, err = queries.GetContentFromPath(context.Background(), childPath)
			// Does not exist in db
			if err != nil {
				_, seasonNo, episodeNo, err := getShowInformation(currentPath, file.Name())
				if err == nil && parentLibrary.ExternalProviderID.Valid {
					episodeInformation, err := client.GetEpisodeInformation(int(parentLibrary.ExternalProviderID.Int64), seasonNo, episodeNo)
					if err == nil {
						queries.AddNewContentFile(context.Background(), storage.AddNewContentFileParams{
							MediaLibraryID:     mediaLibrary.ID,
							CreatedAt:          time.Now(),
							FilePath:           childPath,
							Name:               file.Name(),
							MediaTitle:         episodeInformation.Name,
							Description:        sql.NullString{String: episodeInformation.Overview, Valid: true},
							CoverUrl:           sql.NullString{String: "https://image.tmdb.org/t/p/w500/" + episodeInformation.PosterPath, Valid: true},
							ParentID:           sql.NullInt64{Int64: parentLibrary.ID, Valid: true},
							ExternalProvider:   sql.NullString{String: "tmdb", Valid: true},
							ExternalProviderID: sql.NullInt64{Int64: int64(episodeInformation.Id), Valid: true},
							MediaType:          "episode",
							Classifier:         fmt.Sprintf("S%dE%d", seasonNo, episodeNo),
						})
						continue
					}
				}
				_, err = queries.AddNewContentFile(context.Background(), storage.AddNewContentFileParams{
					MediaLibraryID: mediaLibrary.ID,
					CreatedAt:      time.Now(),
					FilePath:       childPath,
					Name:           file.Name(),
					MediaTitle:     file.Name(),
					ParentID:       sql.NullInt64{Int64: parentLibrary.ID, Valid: true},
					MediaType:      "episode",
					Classifier:     fmt.Sprintf("S%dE%d", seasonNo, episodeNo),
				})
			}
		}
	}

	return nil
}
