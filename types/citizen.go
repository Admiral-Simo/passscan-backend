package types

import (
	"time"
)

type Person struct {
	CNIE                 string    `json:"cin" gorm:"primaryKey;column:cnie"`
	FirstName            string    `json:"first_name" gorm:"column:first_name"`
	LastName             string    `json:"last_name" gorm:"column:last_name"`
	City                 string    `json:"city" gorm:"column:city"`
	Nationality          string    `json:"nationality" gorm:"column:nationality"`
	BirthDate            time.Time `json:"birth_date" gorm:"column:birth_date"`
	ExpireDate           time.Time `json:"expire_date" gorm:"column:expire_date"`
	PossibleNamesAddress []string  `gorm:"type:text[]" json:"possible_names_address"`
}
