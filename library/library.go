package library

import (
	"context"
	"ocelot/auth"
	"ocelot/storage"
	"time"

	"github.com/labstack/echo/v4"
)

func AddLibraryFolder(c echo.Context, u auth.User) error {
	var request struct {
		Path        string `json:"path"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Description string `json:"description"`
	}
	err := c.Bind(&request)
	if err != nil {
		return c.String(500, err.Error())
	}
	if request.Path == "" {
		return c.String(500, "Request is invalid: Path not mentioned.")
	}
	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return c.String(500, err.Error())
	}
	library, err := queries.CreateMediaLibrary(context.Background(), storage.CreateMediaLibraryParams{
		OwnerID:     u.ID,
		Name:        request.Name,
		DevicePath:  request.Path,
		MediaType:   request.Type,
		CreatedAt:   time.Now(),
		Description: request.Description,
	})
	if err != nil {
		return c.String(500, err.Error())
	}

	go ScanLibrary(library.ID)

	return c.NoContent(201)
}
