package db

import (
	"passport_card_analyser/types"
)

func (dba *Adapter) CreateDocument(document types.Document) error {
	err := dba.db.Create(&document).Error
	return err
}

func (dba *Adapter) GetDocuments() ([]*types.Document, error) {
	var documents []*types.Document
	err := dba.db.Find(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (dba *Adapter) GetDocument(docNumber string) (*types.Document, error) {
	var document types.Document
	err := dba.db.Where("document_number = ?", docNumber).First(&document).Error
	if err != nil {
		return nil, err
	}
	return &document, nil
}
