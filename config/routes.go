package config

import "github.com/labstack/echo/v4"

func GetServerInformation(c echo.Context) error {
	serverInformation, err := getServerInformation()
	if err != nil {
		return err
	}
	return c.JSON(200, serverInformation)
}
