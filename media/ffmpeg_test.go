package media

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFfmpeg(t *testing.T) {

	os.Mkdir(os.Getenv("VIDEO_DEST"), 0777)
	prevDir, _ := os.ReadDir(os.Getenv("VIDEO_DEST"))
	prevLen := len(prevDir)
	f := NewFfmpeg("veryfast")

	go f.Start(0, os.Getenv("VIDEO_SRC"), os.Getenv("VIDEO_DEST")+"/stream.m3u8")
	time.Sleep(time.Second * 2)

	go f.Stop()

	newDir, _ := os.ReadDir(os.Getenv("VIDEO_DEST"))

	if prevLen >= len(newDir) {
		fmt.Printf("No files were created.\n")
		t.FailNow()

	}

	os.RemoveAll(os.Getenv("VIDEO_DEST"))

}
