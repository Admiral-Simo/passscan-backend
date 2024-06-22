package ocrscanner

import (
	"fmt"
	"passport_card_analyser/types"
	"time"
	"unicode"
)

func containsTwoSlashes(s string) bool {
	count := 0
	for _, r := range s {
		if r == '/' {
			count++
		}
	}
	return count == 2
}

func containsTwoDots(s string) bool {
	count := 0
	for _, r := range s {
		if r == '.' {
			count++
		}
	}
	return count == 2
}

func containsCINLengthNumbers(s string) bool {
	count := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			count++
		}
	}
	return !unicode.IsDigit(rune(s[0])) && count >= 4
}

func containsTwoSpaces(s string) bool {
	count := 0
	for _, r := range s {
		if r == ' ' {
			count++
		}
	}
	return count == 2
}

func allDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func allUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsUpper(r) && r != ' ' {
			return false
		}
	}
	return true
}

func printArrayString(label string, strings []string) {
	fmt.Printf("%s: { ", label)
	for i, word := range strings {
		fmt.Printf("%d : %s, ", i, word)
	}
	fmt.Printf(" }\n")
}

func printPerson(person *types.Person) {
	printPersonHelperForStrings("CIN:", person.CIN)
	printPersonHelperForStrings("FirstName:", person.FirstName)
	printPersonHelperForStrings("LastName:", person.LastName)
	printPersonHelperForStrings("City:", person.City)
	printPersonHelperForTime("BirthDate:", person.BirthDate)
	printPersonHelperForTime("ExpireDate:", person.ExpireDate)
}

func printPersonHelperForStrings(label, value string) {
	// if string check if empty if int check
	if value == "" {
		fmt.Println(label, "-- none --")
	} else {
		fmt.Println(label, value)
	}
}

func printPersonHelperForTime(label string, value time.Time) {
	// if string check if empty if int check
	if value.IsZero() {
		fmt.Println(label, "-- none --")
	} else {
		fmt.Println(label, value)
	}
}
