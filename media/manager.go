package media

import (
	"errors"
	"ocelot/config"
	"os/exec"
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
