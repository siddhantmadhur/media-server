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
	// Creates the right m3u url for the playback client. i.e. what time to resume, subtitles to use etc.
	e.GET("/media/playback/:mediaId/playback", streamer.GetPlaybackInfo)

	// Once the m3u8 url is made it will be in the format below
	e.GET("/media/content/:mediaId/:streamId/master.m3u8", streamer.GetPlaylistFile)

	// The m3u8 will refer to segments from below
	e.GET("/media/content/:mediaId/segment/:segmentId", streamer.GetLiveStream)
}
