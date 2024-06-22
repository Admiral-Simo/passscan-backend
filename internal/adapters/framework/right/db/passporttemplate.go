package db

import "passport_card_analyser/types"

func (dba *Adapter) CreateTemplate(template types.OCRTemplate) error {
	err := dba.db.Create(&template).Error
	return err
}

func (dba *Adapter) GetTemplateByNationality(nationality string) (*types.OCRTemplate, error) {
	var template types.OCRTemplate
	err := dba.db.Where("nationality = ?", nationality).First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}
