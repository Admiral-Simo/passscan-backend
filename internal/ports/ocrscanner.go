package ports

import "passport_card_analyser/types"

type OCRScannerPost interface {
	ParseDocument(image string) (*types.Document, error)
}

type OCRTextExtractor interface {
	GetContent(image string) ([]string, error)
}
