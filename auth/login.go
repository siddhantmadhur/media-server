package auth

import (
	"context"
	"net/http"
	"ocelot/config"
	"ocelot/storage"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context, cfg *config.Config) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.Bind(&request)
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	conn, query, err := storage.GetConn(cfg)
	defer conn.Close()
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	user, err := query.GetUserFromUsername(context.Background(), request.Username)
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		var result = map[string]string{
			"message": "Password does not match",
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	// TODO: Create session/jwt whatever and send to the user

	return c.NoContent(200)
}
