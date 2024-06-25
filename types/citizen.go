package types

import (
	"time"
)

type PassportData struct {
	CNIE        string    `json:"cin" gorm:"primaryKey;column:cnie"`
	FirstName   string    `json:"first_name" gorm:"column:first_name"`
	LastName    string    `json:"last_name" gorm:"column:last_name"`
	City        string    `json:"city" gorm:"column:city"`
	Sex         string    `json:"sex" gorm:"column:sex"`
	Nationality string    `json:"nationality" gorm:"column:nationality"`
	BirthDate   time.Time `json:"birth_date" gorm:"column:birth_date"`
	ExpireDate  time.Time `json:"expire_date" gorm:"column:expire_date"`
}

type Person struct {
	CNIE                 string         `json:"cin" gorm:"primaryKey;column:cnie"`
	FirstName            string         `json:"first_name" gorm:"column:first_name"`
	LastName             string         `json:"last_name" gorm:"column:last_name"`
	City                 string         `json:"city" gorm:"column:city"`
	Sex                  string         `json:"sex" gorm:"column:sex"`
	Nationality          string         `json:"nationality" gorm:"column:nationality"`
	BirthDate            time.Time      `json:"birth_date" gorm:"column:birth_date"`
	ExpireDate           time.Time      `json:"expire_date" gorm:"column:expire_date"`
	PossibleNamesAddress []PossibleName `json:"possible_names_address" gorm:"foreignKey:PersonCNE;references:CNIE"`
}

// MRZData represents parsed MRZ data
type MRZData struct {
	DocumentType   string `json:"documentType"`
	CountryCode    string `json:"countryCode"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Sex            string `json:"sex"`
	DocumentNumber string `json:"documentNumber"`
	BirthDate      string `json:"birthDate"`
	ExpireDate     string `json:"expireDate"`
	// Add more fields as needed
}

type PossibleName struct {
	ID        uint   `gorm:"primaryKey"`
	PersonCNE string `gorm:"index;column:person_cnie"`
	Name      string `json:"name" gorm:"column:name"`
}
