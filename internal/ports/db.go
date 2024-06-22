package ports

import "passport_card_analyser/types"

type DBPort interface {
	CreatePassport(person types.PersonWithNames) error
	GetPassports() ([]*types.PersonWithNames, error)

	GetPassport(cne string) (*types.PersonWithNames, error)
	CloseDatabase() error
}
