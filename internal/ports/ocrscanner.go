package ports

import "passport_card_analyser/types"

type OCRScannerPost interface {
	ParseCitizen(image string) (*types.Person, []string, error)
}
