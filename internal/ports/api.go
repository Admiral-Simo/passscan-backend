package ports

import (
	"passport_card_analyser/types"
)

type APIPort interface {
	GetPassportData(filepath string, nationality string) (*types.Person, error)
}
