package auth

import (
	"context"
	"ocelot/storage"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.Bind(&request)

	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	err = queries.CreateProfile(context.Background(), storage.CreateProfileParams{
		Username: request.Username,
		Password: request.Password,
		Type:     1,
	})
	if err != nil {
		return err
	}

	return c.NoContent(201)
}
