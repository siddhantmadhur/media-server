package media

import (
	"fmt"
	"log"
	"ocelot/config"
	"os"
	"os/exec"
	"sync"

	"github.com/google/uuid"
)

type Ffmpeg struct {
	lock                  sync.Mutex
	Id                    string
	Proc                  *exec.Cmd
	Preset                string
	CurrentPlaybackSecond int64
	CurrentPath           string
	TranscodePath         string
	MediaId               int64
	StreamUrl             string
}

func NewFfmpeg(preset string, sourcePath string, currentPlaybackSecond int64, c *config.Config, mediaId int64) (*Ffmpeg, error) {
	if preset == "" {
		preset = "veryfast"
	}

	var ffmpeg = Ffmpeg{
		Id:                    uuid.NewString(),
		Preset:                preset,
		CurrentPlaybackSecond: currentPlaybackSecond,
		CurrentPath:           sourcePath,
		MediaId:               mediaId,
	}

	ffmpeg.TranscodePath = fmt.Sprintf("%s/%s", c.CacheDir, ffmpeg.Id)
	ffmpeg.StreamUrl = fmt.Sprintf("/media/%d/streams/%s/master.m3u8", ffmpeg.MediaId, ffmpeg.Id)

	err := os.MkdirAll(ffmpeg.TranscodePath, 0777)

	if err != nil {
		log.Printf("[ERROR]: %s\n", err.Error())
		return &ffmpeg, err
	}

	return &ffmpeg, nil
}

func (f *Ffmpeg) Start() {
	f.Stop()
	f.lock.Lock()
	defer f.lock.Unlock()

	f.Proc = exec.Command("ffmpeg", "-ss", getTimeStamp(f.CurrentPlaybackSecond), "-i", f.CurrentPath, "-preset", f.Preset, "-start_number", fmt.Sprint(f.CurrentPlaybackSecond/2), "-hls_playlist_type", "vod", "-force_key_frames", "expr:gte(t,n_forced*2.0000)", "-hls_time", "2", "-hls_list_size", "0", "-f", "hls", "-y", f.TranscodePath+"/master.m3u8")

	f.Proc.Run()
}

func (f *Ffmpeg) Stop() {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.Proc != nil {
		f.Proc.Process.Kill()
		f.Proc.Wait()
	}

}

func (f *Ffmpeg) SkipTo(segment int64) {
	f.lock.Lock()
	defer f.lock.Unlock()
	if !doesSegmentExist(f, segment) {
		f.Stop()
		f.CurrentPlaybackSecond = segment * 2
		f.Start()
	}
}
