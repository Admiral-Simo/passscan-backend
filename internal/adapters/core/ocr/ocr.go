package ocr

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"passport_card_analyser/internal/adapters/core/types"
	"passport_card_analyser/internal/adapters/core/utilities"
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

func (p *Parser) ParseCitizen() (*types.Person, error) {
	names := []string{}

	lines := strings.Split(p.text, "\n")

	person := &types.Person{}

	var date1, date2 time.Time

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
			return nil, fmt.Errorf("can't parse date %s: %v", line[:10], err)
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

	utilities.PrintArrayString("names", names)

	p.OpenImage()

	for i := 0; i < 3; i++ {
		j := -1
		switch i {
		case 0:
			// this is the names.FirstName
			fmt.Print("give me which index you think is the first name: ")
			fmt.Scan(&j)
			if j < 0 || j > len(names)-1 {
				person.FirstName = ""
				continue
			}
			person.FirstName = names[j]
		case 1:
			// this is the names.LastName
			fmt.Print("give me which index you think is the last name: ")
			fmt.Scan(&j)
			if j < 0 || j > len(names)-1 {
				person.LastName = ""
				continue
			}
			person.LastName = names[j]
		case 2:
			// this is the names.City
			fmt.Print("give me which index you think is the city: ")
			fmt.Scan(&j)
			if j < 0 || j > len(names)-1 {
				person.City = ""
				continue
			}
			person.City = names[j]
		}
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
