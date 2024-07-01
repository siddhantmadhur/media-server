package library

import (
	"context"
	"fmt"
	"ocelot/config"
	"ocelot/storage"
	"strconv"

	"github.com/labstack/echo/v4"
)

// /media/library/content?library=id
func GetContentFromLibrary(c echo.Context, cfg *config.Config) error {
	libraryId, err := strconv.Atoi(c.QueryParam("library"))
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	conn, query, err := storage.GetConn(cfg)
	defer conn.Close()
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	contentFiles, err := query.GetAllContentFiles(context.Background(), int64(libraryId))

	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	var result = []map[string]string{}
	var shows = map[string]([]map[string]string){}

	var contentType string
	var libraryName string
	var libraryDescription string

	for _, content := range contentFiles {
		if content.MediaType.String == "series" {
			contentType = "series"
			libraryName = content.Name_2.String
			libraryDescription = content.Description_2.String
			var show = []map[string]string{}
			show = shows[content.Title]
			show = append(show, map[string]string{
				"id":          fmt.Sprint(content.ID),
				"name":        content.Name,
				"season_no":   fmt.Sprint(content.SeasonNo.Int64),
				"episode_no":  fmt.Sprint(content.EpisodeNo.Int64),
				"title":       content.Title,
				"description": content.Description.String,
				"imdb_id":     fmt.Sprint(content.ImdbID.Int64),
				"media_type":  content.MediaType.String,
				"matched":     fmt.Sprint(content.ImdbID.Valid),
				"extension":   content.Extension,
				"cover_url":   content.CoverUrl.String,
			})
			shows[content.Title] = show
		}
	}

	if contentType == "series" {
		return c.JSON(200, map[string]any{
			"library_name":        libraryName,
			"library_id":          fmt.Sprint(libraryId),
			"library_description": libraryDescription,
			"data":                shows,
		})
	}

	return c.JSON(200, result)
}
