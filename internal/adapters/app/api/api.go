package api

import (
	"passport_card_analyser/internal/ports"
	"passport_card_analyser/types"
)

type Adapter struct {
	ocrscanner ports.OCRScannerPost
}

func NewAdapter(ocrscanner ports.OCRScannerPost) *Adapter {
	return &Adapter{
		ocrscanner: ocrscanner,
	}
}

func (apia Adapter) GetPassportData(filepath string) (*types.Person, []string, error) {
	person, names, err := apia.ocrscanner.ParseCitizen(filepath)
	return person, names, err
}
