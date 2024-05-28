package media

import (
	"context"
	"fmt"
	"log"
	"math"
	"ocelot/config"
	"ocelot/storage"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type MediaManager struct {
	UserSession       string
	RequestedTick     int64
	BufferedUntilTick int64
	FFMPEGProcess     *exec.Cmd
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
	log.Printf("Url: %s\n", c.Request().RequestURI)
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
		log.Printf("Segment Id: %s, Media Id: %d", segment, mediaId)
		if err != nil {
			return c.String(500, err.Error())
		}

		timestamp := time.Duration(time.Second * time.Duration(segmentNo) * 5)
		hours := math.Floor(timestamp.Hours())
		minutes := math.Floor(timestamp.Minutes()) - (hours * 60)
		seconds := timestamp.Seconds() - (hours * 3600) - (minutes * 60)
		formatTimestamp := fmt.Sprintf("%.2d:%.2d:%.2d", int(hours), int(minutes), int(seconds))
		log.Printf("Timestamp: %s\n", formatTimestamp)
		if m.FFMPEGProcess != nil {
			m.FFMPEGProcess.Process.Kill()
			m.FFMPEGProcess = nil
		}
		ffmpeg := exec.Command("ffmpeg", "-ss", formatTimestamp, "-i", m.ContentPath, "-preset", "veryfast", "-start_number", segment, "-hls_playlist_type", "vod", "-force_key_frames", "expr:gte(t,n_forced*5.00)", "-hls_time", "5", "-hls_list_size", "0", "-f", "hls", "-y", fmt.Sprintf("%s/%d-stream.m3u8", m.Config.CacheDir, mediaId))

		ffmpeg.Start()

		m.FFMPEGProcess = ffmpeg

	}

	time.Sleep(time.Second * 1)

	return c.File(fmt.Sprintf("%s/%d-stream%s.ts", m.Config.CacheDir, mediaId, num.FindString(c.Param("segmentId"))))
}
