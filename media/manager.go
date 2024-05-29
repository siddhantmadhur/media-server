package media

import (
	"context"
	"fmt"
	"ocelot/config"
	"ocelot/storage"
	"os"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MediaManager struct {
	UserSession       string
	RequestedTick     int64
	BufferedUntilTick int64
	FFMPEGProcess     *Ffmpeg
	Config            *config.Config
	GeneratingFrame   int
	ContentId         int
	ContentPath       string
}

// TODO: use path and not media id
func (m *MediaManager) GetPlaylistFile(c echo.Context) error {
	num, err := regexp.Compile("[0-9]+")
	mediaId, err := strconv.Atoi(num.FindString(c.Param("mediaId")))

	if err != nil {
		return c.String(500, err.Error())
	}
	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return c.String(500, err.Error())
	}
	media, err := queries.GetContentInfo(context.Background(), int64(mediaId))
	playlist, err := CreatePlaylistHLSFile(media.FilePath, mediaId)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.String(200, playlist)
}

func (m *MediaManager) GetLiveStream(c echo.Context) error {
	mediaId, err := strconv.Atoi(c.Param("mediaId"))
	if m.ContentId != mediaId {
		conn, queries, err := storage.GetConn()
		defer conn.Close()
		if err != nil {
			return c.String(500, err.Error())
		}
		media, err := queries.GetContentInfo(context.Background(), int64(mediaId))
		m.ContentPath = media.FilePath
		m.ContentId = mediaId
	}
	if err != nil {
		return c.String(500, err.Error())
	}
	num, _ := regexp.Compile("[0-9]+")
	_, err = os.ReadFile(fmt.Sprintf("%s/%d-stream%s.ts", m.Config.CacheDir, mediaId, num.FindString(c.Param("segmentId"))))

	if err != nil {
		segment := num.FindString(c.Param("segmentId"))

		segmentNo, err := strconv.Atoi(segment)
		if err != nil {
			return c.String(500, err.Error())
		}

		m.FFMPEGProcess.Start(int64(segmentNo)*3, m.ContentPath, fmt.Sprintf("%s/%d-stream.m3u8", m.Config.CacheDir, mediaId))
		for {
			_, err = os.ReadFile(fmt.Sprintf("%s/%d-stream%s.ts", m.Config.CacheDir, mediaId, num.FindString(c.Param("segmentId"))))
			if err == nil {
				break
			}
		}
	}

	c.Response().Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	return c.File(fmt.Sprintf("%s/%d-stream%s.ts", m.Config.CacheDir, mediaId, num.FindString(c.Param("segmentId"))))
}
