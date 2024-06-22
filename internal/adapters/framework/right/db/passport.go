package db

import (
	"passport_card_analyser/types"
)

func (dba *Adapter) CreatePassport(person types.Person) error {
	err := dba.db.Create(&person).Error
	return err
}

func (dba *Adapter) GetPassports() ([]*types.Person, error) {
	var persons []*types.Person
	err := dba.db.Find(&persons).Error
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (dba *Adapter) GetPassport(cin string) (*types.Person, error) {
	var person types.Person
	err := dba.db.Where("person_cin = ?", cin).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}
