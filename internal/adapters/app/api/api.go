package api

import (
	"fmt"
	"passport_card_analyser/internal/ports"
	"passport_card_analyser/types"
	"strings"
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
	nationality = strings.ToUpper(nationality)
	template, err := apia.database.GetTemplateByNationality(nationality)
	if err != nil {
		return nil, fmt.Errorf("unable to identify nationality %s", nationality)
	}
	_ = template
	// later make the ParseCitizen take the bounds as an input to get the exact data
	person, err := apia.ocrscanner.ParseCitizen(filepath, template.Bounds)
	if err == nil {
		apia.database.CreatePassport(*person)
	}
	return person, err
}
func (apia Adapter) GetTempateNationalities() ([]string, error) {
	return apia.database.GetTemplateNationalities()
}
