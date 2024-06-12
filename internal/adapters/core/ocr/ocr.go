package ocr

import (
	"fmt"
	"io"
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
	cmd := exec.Command("tesseract", p.image, "output", "--psm", "11")
	err := cmd.Run()
	if err != nil {
		return err
	}

	file, err := os.Open("output.txt")

	out, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	os.Remove("output.txt")

	p.text = string(out)
	return nil
}

const (
	slashLayout = "02/01/2006"
	dotLayout   = "02.01.2006"
	minCNE      = 7
	maxCNE      = 8
)

func (p *Parser) ParseCitizen() (*types.Person, error) {
	fmt.Println(p.String())
	re := regexp.MustCompile("[ \n]+")

	parts := re.Split(p.text, -1)

	person := &types.Person{}

	var date1, date2 time.Time

	for _, word := range parts {
		// check for CNE
		if len(word) >= minCNE && len(word) <= maxCNE && utilities.ContainsCNELengthNumbers(word) {
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

func (p *Parser) String() string {
	return p.text
}
