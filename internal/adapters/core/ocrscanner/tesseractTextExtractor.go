package ocrscanner

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type TesseractTextExtractor struct {
}

func (t TesseractTextExtractor) GetContent(image string) ([]string, error) {
	var results []string

	// Helper function to run Tesseract OCR
	runTesseract := func(image string) (string, error) {
		cmd := exec.Command("tesseract", image, "output", "--psm", "6")
		err := cmd.Run()
		if err != nil {
			return "", err
		}

		file, err := os.Open("output.txt")
		if err != nil {
			return "", err
		}
		defer file.Close()

		out, err := io.ReadAll(file)
		if err != nil {
			return "", err
		}

		// Clean up the output file
		os.Remove("output.txt")
		return string(out), nil
	}

	// Helper function to flip the image
	flipImage := func(image string, angle int) error {
		cmd := exec.Command("convert", image, "-rotate", fmt.Sprintf("%d", angle), image)
		return cmd.Run()
	}

	// Process the image for each rotation
	for i := 0; i < 4; i++ {
		// Run Tesseract and capture the result
		result, err := runTesseract(image)
		if err != nil {
			return nil, err
		}
		results = append(results, result)

		// Flip the image for the next iteration
		if i != 3 {
			if err := flipImage(image, 90); err != nil {
				return nil, err
			}
		}
	}

	return results, nil
}
