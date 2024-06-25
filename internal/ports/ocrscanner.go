package ports

import "passport_card_analyser/types"

type OCRScannerPost interface {
	ParsePassport(image string) (*types.MRZData, error)
	ParseIDCard(image string) (*types.Person, error)
}
