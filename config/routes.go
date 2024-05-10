package config

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func IsFinishedSetup(c echo.Context) error {
	finished, err := finishedSetup()
	if err != nil {
		return err
	}
	return c.String(200, fmt.Sprintf("%t", finished))
}

func CreateAdminUser(c echo.Context) error {
	var request struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	err := c.Bind(&request)
	if err != nil {
		return err
	}
	err = createAdminUser(request.Username, request.Password, request.ConfirmPassword)
	type response struct {
		Data  any `json:"data"`
		Error any `json:"error"`
	}
	if err != nil {
		return c.JSON(500, response{
			Error: err.Error(),
		})
	}
	return c.JSON(201, response{
		Data: request,
	})

}
