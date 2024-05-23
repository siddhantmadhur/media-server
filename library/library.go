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
		Path string `json:"path"`
		Name string `json:"name"`
		Type string `json:"type"`
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
	if err != nil {
		return c.String(500, err.Error())
	}
	err = queries.InsertIntoLibrary(context.Background(), storage.InsertIntoLibraryParams{
		Owner:       u.ID,
		Name:        request.Name,
		Path:        request.Path,
		Type:        request.Name,
		CreatedAt:   time.Now(),
		ContentHash: "",
	})

	return c.NoContent(201)
}
