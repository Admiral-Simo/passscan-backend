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

func (apia Adapter) GetPassportData(filepath string) (*types.Person, []string, error) {
	person, names, err := apia.ocrscanner.ParseCitizen(filepath)
	personInfo := types.PersonWithNames{
		Person:               person,
		PossibleNamesAddress: names,
	}
	if err == nil {
		apia.database.CreatePassport(personInfo)
	}
	return person, names, err
}
