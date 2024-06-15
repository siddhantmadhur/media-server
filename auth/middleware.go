package auth

import (
	"ocelot/config"

	"github.com/labstack/echo/v4"
)

type authenticatedRoute func(echo.Context, *User, *config.Config) error

func AuthenticateRoute(next authenticatedRoute, cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c, nil, cfg)
	}
}
