package utilities

import (
	"fmt"
	"passport_card_analyser/internal/adapters/core/types"
	"unicode"
)

func ContainsTwoSlashes(s string) bool {
	count := 0
	for _, r := range s {
		if r == '/' {
			count++
		}
	}
	return count == 2
}

func ContainsTwoDots(s string) bool {
	count := 0
	for _, r := range s {
		if r == '.' {
			count++
		}
	}
	return count == 2
}

func ContainsCNELengthNumbers(s string) bool {
	count := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			count++
		}
	}
	return count >= 4
}

func PrintPerson(person *types.Person) {
	fmt.Println("CNE:", person.CNE)
	fmt.Println("FirstName:", person.FirstName)
	fmt.Println("LastName:", person.LastName)
	fmt.Println("BirthDate:", person.BirthDate)
	fmt.Println("ExpireDate:", person.ExpireDate)
}
