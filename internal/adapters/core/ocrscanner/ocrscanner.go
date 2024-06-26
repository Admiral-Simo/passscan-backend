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

func (ocra Adapter) ParseDocument(image string) (*types.Document, error) {
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

func containsMultipleLessThan(line string) bool {
	count := 0
	for _, r := range line {
		if r == '<' {
			count++
		}
	}
	return count > 5
}
