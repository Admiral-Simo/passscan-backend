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
	CreatePassport(person types.MRZData) error
	GetPassports() ([]*types.MRZData, error)

	GetPassport(cin string) (*types.MRZData, error)
}
