package media

import (
	"errors"
	"ocelot/config"
	"os/exec"

	"github.com/labstack/echo/v4"
)

type MediaManager struct {
	UserSession       string
	ContentFilePath   string
	RequestedTick     int64
	BufferedUntilTick int64
	FFMPEGProcess     *exec.Cmd
	Config            *config.Config
}

// returns the path for the generated stream .m3u8 file or an error
func (m *MediaManager) GenerateM3UFile(mediaId string) (string, error) {
	if mediaId == "" {
		return "", errors.New("Media ID is invalid or empty")
	}

	file := "#EXTM3U\n"
	file += ""

	return "", nil
}

func (m *MediaManager) RestartFFMPEG() error {
	if m.FFMPEGProcess != nil {
		// TODO: Add stop ffmpeg process function here
	}

	m.FFMPEGProcess = exec.Command("echo", "server")

	err := m.FFMPEGProcess.Start()

	return err
}

// TODO: use path and not media id
func (m *MediaManager) GetPlaylistFile(c echo.Context) error {
	mediaId := c.Param("mediaId")
	playlist, err := CreatePlaylistHLSFile(mediaId)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.String(200, playlist)
}
