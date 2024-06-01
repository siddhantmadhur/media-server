package main

import (
	"log"
	"ocelot/auth"
	"ocelot/config"
	"ocelot/library"
	"ocelot/media"
	"ocelot/wizard"
	"os"

	"github.com/labstack/echo/v4"
)

func handler(e *echo.Echo) {

	e.Use(Logger)

	//var ffmpeg = media.NewFfmpeg("veryfast")

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

	var config config.Config
	config.Read()
	streamer, err := media.NewManager(&config)
	if err != nil {
		log.Printf("[Error]: %s\n", err.Error())
		os.Exit(1)
	}
	//streamer.FFMPEGProcess = &ffmpeg
	// Creates the right m3u url for the playback client. i.e. what time to resume, subtitles to use etc.
	e.POST("/media/:mediaId/playback/info", streamer.GetPlaybackInfo)

	// Once the m3u8 url is made it will be in the format below
	//e.GET("/media/content/:mediaId/:streamId/master.m3u8", streamer.GetPlaylistFile)

	// The m3u8 will refer to segments from below
	//e.GET("/media/content/:mediaId/:streamId/:segmentId", streamer.GetLiveStream)
}
