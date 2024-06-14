package ocr

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"passport_card_analyser/internal/adapters/core/types"
	"passport_card_analyser/internal/adapters/core/utilities"
	"sort"
	"strings"
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

func (p *Parser) OpenImage() error {
	cmd := exec.Command("open", p.image)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

const (
	slashLayout = "02/01/2006"
	dotLayout   = "02.01.2006"
	spaceLayout = "02 01 2006"
	minCNE      = 7
	maxCNE      = 8
)

func (p *Parser) ParseCitizen() (*types.Person, []string, error) {
	names := []string{}

	lines := strings.Split(p.text, "\n")

	person := &types.Person{}

	var dates []time.Time

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// check for names
		if utilities.AllUpper(line) {
			names = append(names, line)
		}

		var dateTime time.Time
		var err error

		for _, word := range strings.Split(line, " ") {
			// check for CNE
			if len(word) >= minCNE && len(word) <= maxCNE && utilities.ContainsCNELengthNumbers(word) {
				person.CNE = word
			}

			// check for dates
			switch {
			case utilities.ContainsTwoSlashes(word) && len(word) == 10:
				dateTime, err = time.Parse(slashLayout, word[:10])
			case utilities.ContainsTwoDots(word) && len(word) == 10:
				dateTime, err = time.Parse(dotLayout, word[:10])
			case utilities.AllDigits(word) && utilities.ContainsTwoSpaces(word) && len(word) == 10:
				dateTime, err = time.Parse(spaceLayout, word[:10])
			}
		}

		if err != nil {
			return nil, nil, fmt.Errorf("can't parse date %s: %v", line[:10], err)
		}

		// Store dates temporarily
		if !dateTime.IsZero() {
			dates = append(dates, dateTime)
		}
	}

	// sort dates
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	// assign birth date and expire date
	if len(dates) == 1 {
		person.BirthDate = dates[0]
	} else if len(dates) == 2 {
		person.BirthDate = dates[0]
		person.ExpireDate = dates[1]
	} else if len(dates) == 2 {
		person.BirthDate = dates[0]
		person.ExpireDate = dates[2]
	}

	return person, names, nil
}

func (p *Parser) SetImage(image string) {
	p.image = image
	p.GetContent()
}

func (p *Parser) String() string {
	return p.text
}
