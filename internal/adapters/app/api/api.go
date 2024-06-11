package api

import "passport_card_analyser/internal/adapters/core/types"

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (adap *Adapter) GetPassports() ([]*types.Person, error) {
	return nil, nil
}

func (adap *Adapter) GetPassport() (*types.Person, error) {
	return nil, nil
}
