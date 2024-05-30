package media

import (
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
)

func TestTimestamp(t *testing.T) {

	expected := []string{"00:06:32", "02:31:30"}
	inputs := []int64{392, 9090}

	for idx, input := range inputs {
		got := getTimeStamp(input)
		if expected[idx] != got {
			fmt.Printf("Got: %s, Expected: %s\n", got, expected[idx])
			t.FailNow()
		}
	}

}
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
	// this is where i have my test footage located
	file := os.Getenv("VIDEO_SRC")

	// lol = (length of file * 0.5) + 4
	content, err := CreatePlaylistHLSFile(file, 0)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}
	size := len(strings.Split(strings.ReplaceAll(content, "#EXT-X-DISCONTINUITY\n", ""), "\n"))
	lengthOfFile, _ := GetLengthOfFile(file)
	if size != int(math.Ceil(lengthOfFile))+6 {
		fmt.Printf("[ERROR]: Expected length: %d, got: %d\n", int(math.Ceil(lengthOfFile))+6, size)
		t.FailNow()
	}
}
