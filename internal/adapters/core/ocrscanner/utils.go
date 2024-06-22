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

func isCIN(word string) bool {
	return len(word) >= minCIN && len(word) <= maxCIN && containsCINLengthNumbers(word)
}

const (
	minCIN = 7
	maxCIN = 8
)
