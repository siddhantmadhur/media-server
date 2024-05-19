package media

import (
	"fmt"
	"os"
	"testing"
)

func TestLengthOfFile(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	file := homedir + "/github/ffmpeg-trial/tmnt.mp4"
	fmt.Printf("[READING] %s\n", file)
	size, err := GetLengthOfFile(file)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}
	fmt.Printf("[SUCCESS]: %f\n", size)
}
