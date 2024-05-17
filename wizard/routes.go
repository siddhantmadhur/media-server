package wizard

import (
	"net/http"
	"ocelot/config"

	"github.com/labstack/echo/v4"
)

func WizardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var config config.Config
		err := config.Read()
		if err != nil {
			return err
		}
		if config.FinishedWizard {
			return c.JSON(http.StatusUnauthorized, struct {
				Msg string `json:"msg"`
			}{
				Msg: "Server is already setup and you no longer have access to this feature",
			})
		}

		return next(c)
	}
}

func GetUser(c echo.Context) error {

	return c.String(200, "User")
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
	err = createFirstUser(request.Username, request.Password, request.ConfirmPassword)
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

func IsFinishedSetup(c echo.Context) error {
	var config config.Config
	err := config.Read()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Msg string `json:"msg"`
		}{
			Msg: "[ERROR]: " + err.Error(),
		})
	}
	return c.JSON(200, struct {
		FinishedSetup bool `json:"finished_setup"`
	}{
		FinishedSetup: config.FinishedWizard,
	})
}
