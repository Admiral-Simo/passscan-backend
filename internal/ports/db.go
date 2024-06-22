package ports

import "passport_card_analyser/types"

type DBPort interface {
	CreatePassport(person types.PersonWithNames) error
	GetPassports() ([]*types.PersonWithNames, error)

	GetPassport(cin string) (*types.PersonWithNames, error)

	CreateTemplate(cin string) (*types.PersonWithNames, error)
	CloseDatabase() error
}
