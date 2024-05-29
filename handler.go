package main

import (
	"ocelot/auth"
	"ocelot/config"
	"ocelot/library"
	"ocelot/media"
	"ocelot/wizard"

	"github.com/labstack/echo/v4"
)

func handler(e *echo.Echo) {

	var ffmpeg = media.NewFfmpeg("veryfast")

	// Wizard routes
	e.GET("/wizard/get-first-user", wizard.WizardMiddleware(wizard.GetUser))

	// Server config
	e.GET("/server/information", config.GetServerInformation)

	// Library
	e.POST("/server/media/library", auth.AuthenticateRoute(library.AddLibraryFolder))

	// Auth routes
	e.POST("/auth/login", auth.Login)
	e.POST("/auth/create/user", auth.CreateUser)
	e.GET("/auth/get-user", auth.AuthenticateRoute(auth.GetUserInformation))

	// Streaming routes

	var streamer media.MediaManager
	var config config.Config
	config.Read()
	streamer.Config = &config
	streamer.FFMPEGProcess = &ffmpeg
	e.GET("/media/content/:mediaId", streamer.GetPlaylistFile)
	e.GET("/media/:mediaId/segment/:segmentId", streamer.GetLiveStream)
}
