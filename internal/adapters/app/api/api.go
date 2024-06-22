package api

import (
	"fmt"
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

func (apia Adapter) GetPassportData(filepath string, nationality string) (*types.Person, error) {
	template, err := apia.database.GetTemplateByNationality(nationality)
	if err != nil {
		return nil, fmt.Errorf("unable to identify nationality %s", nationality)
	}
	_ = template
	// later make the ParseCitizen take the bounds as an input to get the exact data
	person, err := apia.ocrscanner.ParseCitizen(filepath)
	if err == nil {
		apia.database.CreatePassport(*person)
	}
	return person, err
}
