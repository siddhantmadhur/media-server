package library

import (
	"context"
	"errors"
	"ocelot/config"
	"ocelot/content"
	"ocelot/content/tmdb"
	"ocelot/content/types"
	"ocelot/storage"
	"os"
	"strings"
)

func MatchContentLibraryWithMetadata(contentId int64, cfg *config.Config) error {

	conn, queries, err := storage.GetConn(cfg)
	defer conn.Close()

	if err != nil {
		return err
	}

	mediaInfo, err := queries.GetContentInfo(context.Background(), contentId)

	if err != nil {
		return err
	}

	relativePath := strings.ReplaceAll(mediaInfo.FilePath, mediaInfo.DevicePath.String, "")

	client, err := content.NewClient(tmdb.Client{
		ApiKey: os.Getenv("TMDB_API_KEY"),
	})
	if err != nil {
		return err
	}

	if mediaInfo.MediaType.String == "shows" {
		shows, err := client.SearchShows(types.SearchParam{
			Query: relativePath,
		})
		if err != nil {
			return err
		}

		if len(shows.Results) == 0 {
			return errors.New("Could not find any results")
		}

	}

	return nil
}
