package processor

import (
	"fmt"
	"os/exec"
)

func Process(input string) {
	wavPath, err := ConvertToWav(input)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}

	outputDir := "transcripts"
	model := "base"

	cmd := exec.Command("whisper", wavPath, "--model", model, "--output_dir", outputDir)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Whisper failed:", err)
	} else {
		fmt.Println("Transcript complete for", input)
	}
}
