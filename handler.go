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

func handler(e *echo.Echo, cfg *config.Config) {

	e.Use(Logger)

	//var ffmpeg = media.NewFfmpeg("veryfast")

	// Wizard routes
	e.GET("/wizard/get-first-user", wizard.WizardMiddleware(wizard.GetUser))

	// Server config
	e.GET("/server/information", cfg.GetServerInformation)
	e.GET("/server/information/folders", auth.AuthenticateOrWizard(library.GetPathFolders, cfg))
	e.POST("/server/information/wizard", cfg.Route(wizard.FinishWizard))

	// Library
	e.POST("/server/media/library", auth.AuthenticateOrWizard(library.AddLibraryFolder, cfg))
	e.GET("/server/media/library", auth.AuthenticateOrWizard(library.GetLibraryFolders, cfg))

	// Auth routes
	e.POST("/auth/create/user", auth.AuthenticateOrWizard(auth.CreateNewUserRoute, cfg))
	e.POST("/auth/login", cfg.Route(auth.Login))
	e.GET("/auth/user", auth.AuthenticateRoute(auth.GetUser, cfg))

	// Streaming routes
	streamer, err := media.NewManager(cfg)
	if err != nil {
		log.Printf("[ERROR]: %s\n", err.Error())
		os.Exit(1)
	}

	// Creates the right m3u url for the playback client. i.e. what time to resume, subtitles to use etc.
	e.POST("/media/:mediaId/playback/info", auth.AuthenticateRoute(streamer.GetPlaybackInfo, cfg))

	// Once the m3u8 url is made it will be in the format below
	e.GET("/media/:mediaId/streams/:sessionId/master.m3u8", streamer.GetMasterPlaylist)

	// The m3u8 will refer to segments from below
	// /media/:mediaId/streams/:sessionId/:segment/stream.ts
	e.GET("/media/:mediaId/streams/:sessionId/:segment/stream.ts", streamer.GetStreamFile)

	e.GET("/media/:mediaId/direct/:fileName", cfg.Route(streamer.GetDirectPlayVideo))

	e.GET("/server/streaming/sessions", streamer.GetAllSessions)
}
