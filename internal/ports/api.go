package ports

import (
	"passport_card_analyser/types"
)

type APIPort interface {
	GetDocumentData(filepath string) (*types.Document, error)
}
