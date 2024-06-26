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
	CreateDocument(person types.Document) error
	GetDocuments() ([]*types.Document, error)

	GetDocument(docNumber string) (*types.Document, error)
}
