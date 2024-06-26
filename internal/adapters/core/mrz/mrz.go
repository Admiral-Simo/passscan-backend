package mrz

import (
	"errors"
	"fmt"
	"passport_card_analyser/types"
	"strconv"
	"strings"
	"unicode"
)

// ParseMRZ parses the given MRZ text and returns Document and an error if any
func ParseMRZ(mrzText string) (*types.Document, error) {
	lines := strings.Split(mrzText, "\n")

	// Check number of lines and handle different MRZ formats (TD1, TD2, TD3)
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\u003c", "<")
	}

	switch {
	case len(lines) < 2:
		return &types.Document{}, errors.New("Invalid MRZ data: less than 2 lines")
	case len(lines) == 2:
		return parsePassport(lines[0], lines[1])
	case len(lines) == 3:
		return parseIdCard(lines[0], lines[1], lines[2])
	default:
		return &types.Document{}, errors.New("Unknown MRZ format")
	}
}

// Helper function to parse TD3 format
func parsePassport(line1, line2 string) (*types.Document, error) {
	// Example fields for TD3 format
	documentType := myTrim(line1[0:2])
	countryCode := myTrim(line1[2:5])
	names := line1[5:]
	lastName, firstName := getNames(names)
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

	// Create Document struct
	mrzData := types.Document{
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

func parseIdCard(line1, line2, line3 string) (*types.Document, error) {
	documentType := myTrim(line1[0:2])
	countryCode := myTrim(line1[2:5])
	lastName, firstName := getNames(line3)
	documentNumber := getCNIE(line1)
	sexIndex := findClosestSex(line2, 7)
	var sex, birthDate, expireDate string
	if sexIndex != -1 {
		sex = string(line2[sexIndex])
		birthDate = stringifyDate(line2[sexIndex-7:sexIndex-1], "birth")
		expireDate = stringifyDate(line2[sexIndex+1:sexIndex+7], "expire")
	}
	return &types.Document{
		DocumentType:   documentType,
		CountryCode:    countryCode,
		FirstName:      firstName,
		DocumentNumber: documentNumber,
		Sex:            sex,
		BirthDate:      birthDate,
		ExpireDate:     expireDate,
		LastName:       lastName,
	}, nil
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

func stringifyDate(text string, dateType string) string {
	if len(text) != 6 {
		return "Invalid input length"
	}

	year := text[:2]
	month := text[2:4]
	day := text[4:]

	fullYear, err := strconv.Atoi(year)
	if err != nil {
		return "Invalid year"
	}

	currentYear := 2024 % 100

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

func getCNIE(text string) string {
	starting := 13
	for i := starting; i < len(text); i++ {
		if !unicode.IsDigit(rune(text[i])) && text[i] != '<' {
			cne := strings.ReplaceAll(text[i:], "<", " ")
			cne = strings.TrimSpace(cne)
			return cne
		}
	}
	return ""
}
