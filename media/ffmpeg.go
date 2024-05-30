package media

import (
	"fmt"
	"os/exec"
	"sync"
)

type Ffmpeg struct {
	Mutex       *sync.Mutex
	Process     *exec.Cmd
	Settings    *Preferences
	LastSegment int
}

type Preferences struct {
	Preset string
}

func NewFfmpeg(preset string) Ffmpeg {
	var mu sync.Mutex
	var ffmpeg = Ffmpeg{
		Settings: &Preferences{
			Preset: preset,
		},
		Mutex:   &mu,
		Process: nil,
	}
	return ffmpeg
}

func (f *Ffmpeg) Start(playback int64, path string, transcodedPath string) {
	f.Stop()
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	f.Process = exec.Command("ffmpeg", "-ss", getTimeStamp(playback), "-i", path, "-preset", f.Settings.Preset, "-start_number", fmt.Sprint(playback/2), "-hls_playlist_type", "vod", "-force_key_frames", "expr:gte(t,n_forced*2.0000)", "-hls_time", "2", "-hls_list_size", "0", "-f", "hls", "-y", transcodedPath)

	err := f.Process.Start()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

}

func (f *Ffmpeg) Stop() {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	if f.Process != nil {
		f.Process.Process.Kill()
		f.Process.Wait()
		f.Process = nil
	}
}
