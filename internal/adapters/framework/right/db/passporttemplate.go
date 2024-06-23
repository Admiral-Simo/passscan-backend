package db

import (
	"errors"
	"passport_card_analyser/types"

	"gorm.io/gorm"
)

func (dba Adapter) CreateTemplate(template types.OCRTemplate) error {
	var existingTemplate types.OCRTemplate
	err := dba.db.Where("nationality = ?", template.Nationality).First(&existingTemplate).Error

	// Check if the template already exists
	if err == nil {
		return errors.New("template with this nationality already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// If there was an error other than record not found, return it
		return err
	}

	// If template does not exist, create a new one
	err = dba.db.Create(&template).Error
	return err
}

func (dba *Adapter) UpdateTemplate(template types.OCRTemplate) error {
	// Find the existing template by nationality
	var existingTemplate types.OCRTemplate
	err := dba.db.Where("nationality = ?", template.Nationality).First(&existingTemplate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("template not found")
		}
		return err
	}

	// Update the template ID to match the existing one
	template.ID = existingTemplate.ID

	// Update the OCRTemplate
	if err := dba.db.Save(&template).Error; err != nil {
		return err
	}

	// Update the Bounds associated with the template
	for _, bound := range template.Bounds {
		bound.TemplateID = template.ID
		if err := dba.db.Save(&bound).Error; err != nil {
			return err
		}
	}

	return nil
}

func (dba Adapter) GetTemplateByNationality(nationality string) (*types.OCRTemplate, error) {
	var template types.OCRTemplate
	err := dba.db.Where("nationality = ?", nationality).Preload("Bounds").First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// GetTemplate fetches all OCRTemplate records along with their associated Bounds
func (dba Adapter) GetTemplates() ([]*types.OCRTemplate, error) {
	var templates []*types.OCRTemplate
	err := dba.db.Preload("Bounds").Find(&templates).Error
	if err != nil {
		return nil, err
	}
	return templates, nil
}
