package config

import "github.com/labstack/echo/v4"

func GetServerInformation(c echo.Context) error {
	serverInformation, err := getServerInformation()
	if err != nil {
		return err
	}
	return c.JSON(200, serverInformation)
}

func (cfg *Config) Route(next RouteWithConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c, cfg)
	}
}

type RouteWithConfig func(echo.Context, *Config) error
