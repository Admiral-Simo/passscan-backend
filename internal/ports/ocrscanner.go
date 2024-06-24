package ports

import "passport_card_analyser/types"

type OCRScannerPost interface {
	ParseCitizen(image string, bounds []types.Rectangle) (*types.Person, error)
}
