package ocrscanner

import (
	"fmt"
	"passport_card_analyser/internal/adapters/core/mrz"
	"passport_card_analyser/types"
	"strings"
	"time"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ocra Adapter) ParsePassport(image string) (*types.MRZData, error) {
	text, err := getContent(image)
	if err != nil {
		return nil, err
	}

	// get all the long words that contain multiple < signs

	mrz_text := []string{}
	text = strings.TrimSpace(text)
	for _, line := range strings.Split(text, "\n") {
		if containsMultipleLessThan(line) {
			mrz_text = append(mrz_text, strings.TrimSpace(line))
		}
	}

	return mrz.ParseMRZ(strings.Join(mrz_text, "\n"))
}

type PassportData struct {
	DocumentType    string
	DocumentSubtype string
	CountryCode     string
	LastName        string
	FirstName       string
	DocumentNumber  string
	Nationality     string
	DateOfBirth     string
	Sex             string
	ExpirationDate  string
	PersonalNumber  string
	OptionalData    string
	Checksum        string
}

func containsMultipleLessThan(line string) bool {
	count := 0
	for _, r := range line {
		if r == '<' {
			count++
		}
	}
	return count > 5
}

func (ocra Adapter) ParseIDCard(image string) (*types.Person, error) {
	text, err := getContent(image)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(text, "\n")
	person := &types.Person{}
	var dates []time.Time
	var names []string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if allUpper(line) {
			names = append(names, line)
		}

		if err := ocra.parseLine(line, person, &dates); err != nil {
			return nil, fmt.Errorf("can't parse line %s: %v", line, err)
		}
	}

	sortDates(dates)
	assignDates(person, dates)

	return person, nil
}

func (p Adapter) parseLine(line string, person *types.Person, dates *[]time.Time) error {
	words := strings.Split(line, " ")

	for _, word := range words {
		if isCNIE(word) {
			person.CNIE = word
		}

		if dateTime, err := parseDate(word); err == nil {
			*dates = append(*dates, dateTime)
		} else if err != errNoDate {
			return err
		}
	}

	return nil
}

func parseDate(word string) (time.Time, error) {
	var layouts = []string{slashLayout, dotLayout, spaceLayout}
	var parsers = []func(string) bool{
		containsTwoSlashes,
		containsTwoDots,
		allDigits,
	}

	for i, parseFunc := range parsers {
		if parseFunc(word) && len(word) == 10 {
			dateTime, err := time.Parse(layouts[i], word[:10])
			if err != nil {
				return time.Time{}, err
			}
			return dateTime, nil
		}
	}

	return time.Time{}, errNoDate
}

const (
	slashLayout = "02/01/2006"
	dotLayout   = "02.01.2006"
	spaceLayout = "02 01 2006"
)
