package types

import "time"

type Person struct {
	CNE        string
	FirstName  string
	LastName   string
	City       string
	BirthDate  time.Time
	ExpireDate time.Time
}

type PersonWithNames struct {
	Person               *Person  `json:"person"`
	PossibleNamesAddress []string `json:"possible_names_address"`
}
