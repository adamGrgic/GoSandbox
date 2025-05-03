package processor

import (
	"io"
	"os"
	"os/exec"
)

func SaveFile(src io.Reader, dstPath string) (string, error) {
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return dstPath, err
}

func ConvertToWav(input string) (string, error) {
	output := input + ".wav"
	cmd := exec.Command("ffmpeg", "-i", input, "-ar", "16000", "-ac", "1", "-y", output)
	return output, cmd.Run()
}
