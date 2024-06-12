package ocr

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"os/exec"
	"passport_card_analyser/internal/adapters/core/types"
	"passport_card_analyser/internal/adapters/core/utilities"
	"regexp"
	"time"
)

type Parser struct {
	image string
	text  string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) GetContent() error {
	var result bytes.Buffer

	cmd := exec.Command("tesseract", p.image, "-", "stdout")
	cmd.Stdout = &result
	err := cmd.Run()
	if err != nil {
		return err
	}

	// execute this command and get the ouput into the result variable
	p.text = result.String()
	return nil
}

func (p *Parser) GetContentOfPortion(x, y, width, height int) error {
	// Open the image file
	file, err := os.Open(p.image)
	if err != nil {
		return fmt.Errorf("error opening image: %w", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding image: %w", err)
	}

	// Define the cropping rectangle
	rect := image.Rect(x, y, x+width, y+height)

	// Create a new image for the cropped portion
	croppedImg := image.NewRGBA(rect)

	// Draw the cropped portion onto the new image
	draw.Draw(croppedImg, rect, img, image.Point{x, y}, draw.Src)

	// Save the cropped image to a temporary file
	tmpFile, err := os.CreateTemp("", "cropped_*.jpg")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}

	err = jpeg.Encode(tmpFile, croppedImg, nil)
	if err != nil {
		return fmt.Errorf("error encoding cropped image: %w", err)
	}

	// Close the temporary file to ensure it is written
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("error closing temporary file: %w", err)
	}

	// Run Tesseract on the cropped image
	var result bytes.Buffer
	cmd := exec.Command("tesseract", tmpFile.Name(), "-", "stdout")
	cmd.Stdout = &result
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running tesseract: %w", err)
	}
	p.text = result.String()
	return nil
}

const (
	slashLayout = "02/01/2006"
	dotLayout   = "02.01.2006"
	minCNE      = 7
	maxCNE      = 8
)

func (p *Parser) ParseCitizen() (*types.Person, error) {
	re := regexp.MustCompile("[ \n]+")

	parts := re.Split(p.text, -1)

	person := &types.Person{}

	var date1, date2 time.Time

	for _, word := range parts {
		// check for CNE
		if len(word) >= minCNE && len(word) <= maxCNE {
			person.CNE = word
		}

		var dateTime time.Time
		var err error

		// check for dates
		switch {
		case utilities.ContainsTwoSlashes(word) && len(word) == 10:
			dateTime, err = time.Parse(slashLayout, word[:10])
		case utilities.ContainsTwoDots(word) && len(word) == 10:
			dateTime, err = time.Parse(dotLayout, word[:10])
		}

		if err != nil {
			return nil, fmt.Errorf("can't parse date %s: %v", word[:10], err)
		}

		// Store dates temporarily
		if !dateTime.IsZero() {
			if date1.IsZero() {
				date1 = dateTime
			} else if date2.IsZero() {
				date2 = dateTime
			}
		}
	}

	// swap if needed
	if !date1.IsZero() && !date1.IsZero() {
		if date1.Before(date2) {
			person.BirthDate = date1
			person.ExpireDate = date2
		} else {
			person.BirthDate = date2
			person.ExpireDate = date1
		}
	} else if !date1.IsZero() {
		person.BirthDate = date1
	} else if !date2.IsZero() {
		person.BirthDate = date1
	}

	return person, nil
}

func (p *Parser) SetImage(image string) {
	p.image = image
	p.GetContent()
}
