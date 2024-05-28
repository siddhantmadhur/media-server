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
	ContentFilePath   string
	RequestedTick     int64
	BufferedUntilTick int64
	FFMPEGProcess     *exec.Cmd
	Config            *config.Config
	GeneratingFrame   int
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

func (m *MediaManager) GetSegmentFile(c echo.Context) error {
	log.Printf("Url: %s\n", c.Request().RequestURI)
	mediaId, err := strconv.Atoi(c.Param("mediaId"))
	if err != nil {
		return c.String(500, err.Error())
	}
	num, _ := regexp.Compile("[0-9]+")
	_, err = os.Open(fmt.Sprintf("%s/%d-stream%s.ts", m.Config.CacheDir, mediaId, num.FindString(c.Param("segmentId"))))
	if err != nil {
		log.Printf("File %s not generated!\n", fmt.Sprintf("%s/%d-stream%s.ts", m.Config.CacheDir, mediaId, num.FindString(c.Param("segmentId"))))
		conn, queries, err := storage.GetConn()
		defer conn.Close()
		if err != nil {
			return c.String(500, err.Error())
		}
		media, err := queries.GetContentInfo(context.Background(), int64(mediaId))
		if err != nil {
			return c.String(500, err.Error())
		}
		segment := num.FindString(c.Param("segmentId"))

		segmentNo, err := strconv.Atoi(segment)
		if err != nil {
			return c.String(500, err.Error())
		}

		timestamp := time.Duration(time.Second * time.Duration(segmentNo) * 2)
		hours := math.Floor(timestamp.Hours())
		minutes := math.Floor(timestamp.Minutes()) - (hours * 60)
		seconds := timestamp.Seconds() - (hours * 3600) - (minutes * 60)
		formatTimestamp := fmt.Sprintf("%.2d:%.2d:%.2d", int(hours), int(minutes), int(seconds))
		log.Printf("Timestamp: %s\n", formatTimestamp)
		if m.FFMPEGProcess != nil {
			m.FFMPEGProcess.Process.Kill()
			m.FFMPEGProcess.Wait()
		}

		m.FFMPEGProcess = exec.Command("ffmpeg", "-ss", formatTimestamp, "-i", media.FilePath, "-preset", "veryfast", "-start_number", segment, "-hls_playlist_type", "vod", "-force_key_frames", "expr:gte(t,n_forced*2.00000)", "-hls_time", "2", "-hls_list_size", "0", "-f", "hls", "-y", fmt.Sprintf("%s/%d-stream.m3u8", m.Config.CacheDir, mediaId))

		m.FFMPEGProcess.Start()
		time.Sleep(time.Second * 2)

	}
	return c.File(fmt.Sprintf("%s/%d-stream%s.ts", m.Config.CacheDir, mediaId, num.FindString(c.Param("segmentId"))))
}
