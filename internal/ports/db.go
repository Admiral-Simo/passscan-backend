package ports

import "passport_card_analyser/types"

type DBPort interface {
	// passports
	dBPassport
	dBPassportTemplate
	dBCleaner
}

type dBCleaner interface {
	CloseDatabase() error
}

type dBPassportTemplate interface {
	CreateTemplate(template types.OCRTemplate) error
	GetTemplateByNationality(nationality string) (*types.OCRTemplate, error)
}

type dBPassport interface {
	CreatePassport(person types.Person) error
	GetPassports() ([]*types.Person, error)

	GetPassport(cin string) (*types.Person, error)
}
