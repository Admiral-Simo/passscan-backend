package ocrscanner

import (
	"fmt"
	"passport_card_analyser/internal/adapters/core/mrz"
	"passport_card_analyser/internal/ports"
	"passport_card_analyser/types"
	"strings"
)

type Adapter struct {
	textExtractor ports.OCRTextExtractor
}

func NewAdapter(textExtractor ports.OCRTextExtractor) *Adapter {
	return &Adapter{
		textExtractor: textExtractor,
	}
}

func (ocra Adapter) ParseDocument(image string) (*types.Document, error) {
	texts, err := ocra.textExtractor.GetContent(image)
	if err != nil {
		return nil, err
	}

	// get all the long words that contain multiple < signs

	for _, s := range texts {
		result, _ := ocra.extractMRZ(s)
		doc, err := mrz.ParseMRZ(result)
		if err != nil {
			continue
		}
		return doc, nil
	}

	return &types.Document{}, fmt.Errorf("not detected passport or id card.")
}

func (ocra Adapter) extractMRZ(s string) (string, error) {

	containsMultipleLessThan := func(line string) bool {
		count := 0
		for _, r := range line {
			if r == '<' {
				count++
			}
		}
		return count > 5
	}

	mrz_text := []string{}
	s = strings.TrimSpace(s)
	for _, line := range strings.Split(s, "\n") {
		if containsMultipleLessThan(line) {
			mrz_text = append(mrz_text, strings.TrimSpace(line))
		}
	}

	return strings.Join(mrz_text, "\n"), nil
}
