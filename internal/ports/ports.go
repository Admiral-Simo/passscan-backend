package ports

import "passport_card_analyser/internal/adapters/core/types"

type APIocr interface {
	GetPassports() ([]*types.Person, error)
	GetPassport() (*types.Person, error)
}
