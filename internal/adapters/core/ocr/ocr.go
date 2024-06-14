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
	minCNE = 7
	maxCNE = 8
)

func (p *Parser) ParseCitizen() (*types.Person, []string, error) {
	lines := strings.Split(p.text, "\n")
	person := &types.Person{}
	var dates []time.Time
	var names []string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if utilities.AllUpper(line) {
			names = append(names, line)
		}

		if err := p.parseLine(line, person, &dates); err != nil {
			return nil, nil, fmt.Errorf("can't parse line %s: %v", line, err)
		}
	}

	utilities.SortDates(dates)
	utilities.AssignDates(person, dates)

	return person, names, nil
}

func (p *Parser) parseLine(line string, person *types.Person, dates *[]time.Time) error {
	words := strings.Split(line, " ")

	for _, word := range words {
		if isCNE(word) {
			person.CNE = word
		}

		if dateTime, err := utilities.ParseDate(word); err == nil {
			*dates = append(*dates, dateTime)
		} else if err != utilities.ErrNoDate {
			return err
		}
	}

	return nil
}

func isCNE(word string) bool {
	return len(word) >= minCNE && len(word) <= maxCNE && utilities.ContainsCNELengthNumbers(word)
}

func (p *Parser) SetImage(image string) {
	p.image = image
	p.GetContent()
}

func (p *Parser) String() string {
	return p.text
}
