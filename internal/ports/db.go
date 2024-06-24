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
	GetTemplateByNationality(nationality string) (*types.OCRTemplate, error)
	GetTemplates() ([]*types.OCRTemplate, error)
	GetTemplateNationalities() ([]string, error)

	CreateTemplate(template types.OCRTemplate) error
	UpdateTemplate(template types.OCRTemplate) error
}

type dBPassport interface {
	CreatePassport(person types.Person) error
	GetPassports() ([]*types.Person, error)

	GetPassport(cin string) (*types.Person, error)
}
