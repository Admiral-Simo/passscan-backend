package api

import (
	"passport_card_analyser/internal/ports"
	"passport_card_analyser/types"
)

type Adapter struct {
	ocrscanner ports.OCRScannerPost
	database   ports.DBPort
}

func NewAdapter(ocrscanner ports.OCRScannerPost, database ports.DBPort) *Adapter {
	return &Adapter{
		ocrscanner: ocrscanner,
		database:   database,
	}
}

func (apia Adapter) GetPassportData(filepath string) (*types.Person, error) {
	// later make the ParseCitizen take the bounds as an input to get the exact data
	person, err := apia.ocrscanner.ParseCitizen(filepath)
	if err == nil {
		apia.database.CreatePassport(*person)
	}
	return person, err
}
