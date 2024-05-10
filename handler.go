package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/siddhantmadhur/media-server/auth"
	"github.com/siddhantmadhur/media-server/config"
)

func handler(e *echo.Echo) {

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health OK!")
	})
	e.GET("/config/finished", config.IsFinishedSetup)
	e.POST("/config/create-admin", config.CreateAdminUser)

	e.POST("/auth/login", auth.Login)
}
