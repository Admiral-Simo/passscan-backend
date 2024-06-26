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

func (apia Adapter) GetDocumentData(filepath string) (*types.Document, error) {
	document, err := apia.ocrscanner.ParseDocument(filepath)
	if err == nil {
		apia.database.CreateDocument(*document)
	}
	return document, err
}
