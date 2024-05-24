package library

import (
	"fmt"
	"os"
	"testing"
)

func TestProbe(t *testing.T) {
	path := os.Getenv("VIDEO_FILE")
	if path == "" {
		fmt.Printf(`[ERROR]: "VIDEO_FILE" env variable not configured.`)
		t.FailNow()
	}
	probe, err := FFprobe(path)
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("[Probe]\nName: %s\nDuration: %s\nProbe Score: %d\n", probe.Format.Filename, probe.Format.Duration, probe.Format.ProbeScore)
}
