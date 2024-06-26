package ports

import (
	"passport_card_analyser/types"
)

type APIPort interface {
	GetPassportData(filepath string) (*types.Document, error)
	GetIDCardData(filepath string) (*types.Document, error)
}
