package media

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetLengthOfFile(path string) (float64, error) {
	probe := exec.Command("ffprobe", "-i", path, "-show_entries", "format=duration", "-v", "quiet", "-of", `csv=p=0`)
	output, err := probe.Output()
	if err != nil {
		return 0, err
	}
	size, err := strconv.ParseFloat(strings.ReplaceAll(string(output), "\n", ""), 64)
	return size, err
}

// Returns content of m3u8 file as a string or error
func CreatePlaylistHLSFile(path string, mediaId string) (string, error) {
	size, err := GetLengthOfFile(path)
	counter := size
	if err != nil {
		return "", err
	}
	content := "#EXTM3U\n"
	content += "#EXT-X-VERSION:3\n"
	content += "#EXT-X-TARGETDURATION:2\n"
	content += "#EXT-X-MEDIA-SEQUENCE:0\n"

	idx := 0
	for counter > 0.0 {
		newTime := 0.0
		if counter >= 2.0 {
			newTime = 2.0
		} else {
			newTime = counter
		}
		content += fmt.Sprintf("#EXTINF:%.6f,\n", newTime)
		content += fmt.Sprintf("/media/playback/%d/stream%d.ts\n", mediaId, idx)
		counter -= newTime
		idx += 1
	}

	content += "#EXT-X-ENDLIST"
	return content, nil

}
