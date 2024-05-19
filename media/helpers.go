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
		fmt.Printf("[ERROR]:\n")
		fmt.Print(string(output))
		fmt.Printf("\n[END OF ERROR]\n")
		return 0, err
	}
	size, err := strconv.ParseFloat(strings.ReplaceAll(string(output), "\n", ""), 64)
	return size, err
}
