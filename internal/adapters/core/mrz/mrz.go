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
	documentType := line1[0:2]
	countryCode := line1[2:5]
	names := line1[5:]
	firstName, lastName := getNames(names)
	documentNumber := line2[0:9]
	sex := string(line2[20])
	// dates
	birthDate := stringifyDate(line2[13:19])
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

// change from YYMMDD to DD/MM/YYYY
func stringifyDate(text string) string {
	// Check if the input string is exactly 6 characters long
	if len(text) != 6 {
		return "Invalid input length"
	}

	// Extract year, month, and day parts from the input string
	year := text[:2]
	month := text[2:4]
	day := text[4:]

	// Convert the year to YYYY format (assuming 2000s)
	fullYear, err := strconv.Atoi(year)
	if err != nil {
		return "Invalid year"
	}
	fullYear += 1900

	// Format the date as DD/MM/YYYY
	formattedDate := fmt.Sprintf("%s/%s/%d", day, month, fullYear)
	return formattedDate
}
