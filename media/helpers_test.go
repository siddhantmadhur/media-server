package media

import (
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
)

func TestLengthOfFile(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	// this is where i have my test footage located
	file := homedir + "/github/ffmpeg-trial/tmnt.mp4"
	_, err := GetLengthOfFile(file)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}
}

func TestPlaylistCreation(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	// this is where i have my test footage located
	file := homedir + "/github/ffmpeg-trial/tmnt.mp4"

	// lol = (length of file * 0.5) + 4
	content, err := CreatePlaylistHLSFile(file)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}
	size := len(strings.Split(content, "\n"))
	lengthOfFile, _ := GetLengthOfFile(file)
	if size != int(math.Ceil(lengthOfFile))+5 {
		fmt.Printf("[ERROR]: Expected length: %d, got: %d\n", int(math.Ceil(lengthOfFile))+5, size)
		t.FailNow()
	}
}
