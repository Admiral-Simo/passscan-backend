package ports

import "passport_card_analyser/types"

type OCRScannerPost interface {
	ParsePassport(image string) (*types.Document, error)
	ParseIDCard(image string) (*types.Document, error)
}
