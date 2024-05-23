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

	// Server config
	e.GET("/server/information", config.GetServerInformation)

	// Auth routes
	e.POST("/auth/login", auth.Login)
	e.POST("/auth/create/user", auth.CreateUser)
	e.GET("/auth/get-user", auth.AuthenticateRoute(auth.GetUserInformation))

}
