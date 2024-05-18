package main

import (
	"ocelot/auth"
	"ocelot/config"
	"ocelot/wizard"

	"github.com/labstack/echo/v4"
)

func handler(e *echo.Echo) {

	// Wizard routes
	e.GET("/wizard/get-first-user", wizard.WizardMiddleware(wizard.GetUser))
	e.GET("/wizard/is-finished", wizard.IsFinishedSetup)
	e.POST("/wizard/create-first-user", wizard.WizardMiddleware(wizard.CreateAdminUser))

	// Server config
	e.GET("/server/information", config.GetServerInformation)

	// Auth routes
	e.POST("/auth/login", auth.Login)
	e.GET("/auth/get-user", auth.AuthenticateRoute(auth.GetUserInformation))
}
