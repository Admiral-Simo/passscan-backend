package types

import "time"

type Person struct {
	CNE        string
	FirstName  string
	LastName   string
	BirthDate  time.Time
	ExpireDate time.Time
}
