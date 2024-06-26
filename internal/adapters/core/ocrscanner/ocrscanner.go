package ocrscanner

import (
	"passport_card_analyser/internal/adapters/core/mrz"
	"passport_card_analyser/types"
	"strings"
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

func (ocra Adapter) ParseIDCard(image string) (*types.MRZData, error) {

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
	// text, err := getContent(image, "id")
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	// lines := strings.Split(text, "\n")
	// person := &types.MRZData{}
	// var dates []time.Time
	// var names []string
	//
	//	for _, line := range lines {
	//		line = strings.TrimSpace(line)
	//
	//		if allUpper(line) {
	//			names = append(names, line)
	//		}
	//
	//		if err := ocra.parseLine(line, person, &dates); err != nil {
	//			return nil, fmt.Errorf("can't parse line %s: %v", line, err)
	//		}
	//	}
	//
	// sortDates(dates)
	// assignDates(person, dates)
	//
	// return person, nil
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
