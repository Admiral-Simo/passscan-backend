package mrz

import (
	"errors"
	"fmt"
	"passport_card_analyser/types"
	"strconv"
	"strings"
)

type DocumentType byte

const (
	TD1 DocumentType = iota
	TD2
	TD3
)

func (d DocumentType) String() string {
	switch d {
	case TD1:
		return "TD1"
	case TD2:
		return "TD2"
	case TD3:
		return "TD3"
	}
	return "UNKNOWN"
}

// ParseMRZ parses the given MRZ text and returns MRZData and an error if any
func ParseMRZ(mrzText string) (*types.MRZData, error) {
	lines := strings.Split(mrzText, "\n")

	// Check number of lines and handle different MRZ formats (TD1, TD2, TD3)
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\u003c", "<")
	}

	switch {
	case len(lines) < 2:
		return &types.MRZData{}, errors.New("Invalid MRZ data: less than 2 lines")
	case len(lines) == 3:
		return parseTD1(lines[0], lines[1], lines[2])
	case len(lines[0]) > 40:
		return parseTD2(lines[0], lines[1])
	default:
		return &types.MRZData{}, errors.New("Unknown MRZ format")
	}
}

// Helper function to parse TD1 format
func parseTD1(line1, line2, line3 string) (*types.MRZData, error) {
	// Example fields for TD1 format
	documentType := line1[0:2]
	countryCode := line1[2:5]
	documentNumber := line1[5:14]
	nationality := line2[15:18]
	sex := string(line2[7])
	names := line3
	firstName, lastName := getNames(names)
	_ = nationality

	// Create MRZData struct
	mrzData := types.MRZData{
		DocumentType:   documentType,
		CountryCode:    countryCode,
		FirstName:      firstName,
		LastName:       lastName,
		DocumentNumber: documentNumber,
		Sex:            sex,
		// Add more fields as needed
	}

	return &mrzData, nil
}

// Helper function to parse TD3 format
func parseTD2(line1, line2 string) (*types.MRZData, error) {
	// Example fields for TD3 format
	documentType := myTrim(line1[0:2])
	countryCode := myTrim(line1[2:5])
	names := line1[5:]
	firstName, lastName := getNames(names)
	documentNumber := myTrim(line2[0:9])
	sexIndex := findClosestSex(line2, 20)
	var sex, birthDate, expireDate string
	if sexIndex != -1 {
		sex = string(line2[sexIndex])
		birthDate = stringifyDate(line2[sexIndex-7:sexIndex-1], "birth")
		expireDate = stringifyDate(line2[sexIndex+1:sexIndex+7], "expire")
	}
	// dates
	names = strings.ReplaceAll(names, "<", " ")

	// Create MRZData struct
	mrzData := types.MRZData{
		DocumentType:   documentType,
		CountryCode:    countryCode,
		FirstName:      firstName,
		LastName:       lastName,
		DocumentNumber: documentNumber,
		Sex:            sex,
		BirthDate:      birthDate,
		ExpireDate:     expireDate,
		// Add more fields as needed
	}

	return &mrzData, nil
}

func getNames(text string) (string, string) {
	parts := strings.Split(text, "<<")

	// Filter out empty strings from the parts
	var filteredParts []string
	for _, part := range parts {
		if part != "" {
			filteredParts = append(filteredParts, part)
		}
	}

	// Assume the first part is the FirstName and the remaining part is the LastName
	var firstName, lastName string
	if len(filteredParts) > 0 {
		firstName = filteredParts[0]
		firstName = strings.ReplaceAll(firstName, "<", " ")
		firstName = strings.TrimSpace(firstName)
	}
	if len(filteredParts) > 1 {
		lastName = strings.Join(filteredParts[1:], " ")
		lastName = strings.ReplaceAll(lastName, "<", " ")
		lastName = strings.TrimSpace(lastName)
	}

	return firstName, lastName
}

// convert from YYMMDD to DD/MM/YYYY
func stringifyDate(text string, dateType string) string {
	// Check if the input string is exactly 6 characters long
	if len(text) != 6 {
		return "Invalid input length"
	}

	// Extract year, month, and day parts from the input string
	year := text[:2]
	month := text[2:4]
	day := text[4:]

	// Convert the year to YYYY format
	fullYear, err := strconv.Atoi(year)
	if err != nil {
		return "Invalid year"
	}

	// Assume the year is in the 1900s if the year is greater than the current year, otherwise 2000s
	currentYear := 2024 % 100 // Last two digits of the current year

	switch dateType {
	case "expire":
		fullYear += 2000
	case "birth":
		if fullYear > currentYear {
			fullYear += 1900
		} else {
			fullYear += 2000
		}
	}

	// Format the date as DD/MM/YYYY
	formattedDate := fmt.Sprintf("%s/%s/%d", day, month, fullYear)
	return formattedDate
}

func myTrim(input string) string {
	input = strings.ReplaceAll(input, "<", " ")
	input = strings.TrimSpace(input)
	return input
}

func findClosestSex(line string, index int) int {
	outOfBounds := func(i, j int) bool {
		return i >= len(line) || j < 0
	}

	i, j := index, index
	for {
		if outOfBounds(i, j) {
			return -1
		}
		if line[i] == 'F' || line[i] == 'M' {
			return i
		}
		if line[j] == 'F' || line[j] == 'M' {
			return j
		}
		i++
		j--
	}
}
