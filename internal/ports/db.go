package ports

import "passport_card_analyser/types"

type DBPort interface {
	// passports
	dBPassport
	dBCleaner
}

type dBCleaner interface {
	CloseDatabase() error
}

type dBPassport interface {
	CreatePassport(person types.Person) error
	GetPassports() ([]*types.Person, error)

	GetPassport(cin string) (*types.Person, error)
}
