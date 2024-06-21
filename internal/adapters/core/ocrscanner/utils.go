package ocrscanner

import (
	"io"
	"os"
	"os/exec"
)

func getContent(image string) (string, error) {
	cmd := exec.Command("tesseract", image, "output", "--psm", "11")
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	file, err := os.Open("output.txt")

	out, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	os.Remove("output.txt")

	return string(out), nil
}

func isCNE(word string) bool {
	return len(word) >= minCNE && len(word) <= maxCNE && containsCNELengthNumbers(word)
}

const (
	minCNE = 7
	maxCNE = 8
)
