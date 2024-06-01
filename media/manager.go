package media

import (
	"context"
	"fmt"
	"ocelot/config"
	"ocelot/storage"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Manager struct {
	Sessions map[string]*Ffmpeg
	Config   *config.Config
}

func NewManager(cfg *config.Config) (*Manager, error) {
	var m Manager
	m.Config = cfg
	m.Sessions = make(map[string]*Ffmpeg)

	return &m, nil
}

// /media/:mediaId/playback/info
func (m *Manager) GetPlaybackInfo(c echo.Context) error {

	var request struct {
		Preset          string `json:"preset"`
		PlaybackSeconds int64  `json:"playback_seconds"`
	}

	c.Bind(&request)

	mediaId, err := strconv.Atoi(c.Param("mediaId"))
	if err != nil {
		return c.String(500, err.Error())
	}

	conn, queries, err := storage.GetConn()
	defer conn.Close()

	if err != nil {
		return c.String(500, err.Error())
	}

	content, err := queries.GetContentInfo(context.Background(), int64(mediaId))
	if err != nil {
		return c.String(500, err.Error())
	}

	ffmpeg, err := NewFfmpeg(request.Preset, content.FilePath, request.PlaybackSeconds, m.Config, int64(mediaId))
	if err != nil {
		return c.String(500, err.Error())
	}

	m.Sessions[ffmpeg.Id] = ffmpeg

	var response = map[string]string{
		"session_id": ffmpeg.Id,
		"media_id":   fmt.Sprint(mediaId),
		"preset":     ffmpeg.Preset,
		"stream_url": ffmpeg.StreamUrl,
	}

	go ffmpeg.Start()

	return c.JSON(201, response)
}
