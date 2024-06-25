package mrz

import (
	"errors"
	"fmt"
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

// MRZData represents parsed MRZ data
type MRZData struct {
	DocumentType   string
	CountryCode    string
	Nationality    string
	Names          string
	Sex            string
	DocumentNumber string
	BirthDate      string
	// Add more fields as needed
}

// ParseMRZ parses the given MRZ text and returns MRZData and an error if any
func ParseMRZ(mrzText string) (MRZData, error) {
	lines := strings.Split(mrzText, "\n")

	// Check number of lines and handle different MRZ formats (TD1, TD2, TD3)
	switch {
	case len(lines) < 2:
		return MRZData{}, errors.New("Invalid MRZ data: less than 2 lines")
	case len(lines) == 3:
		return parseTD1(lines[0], lines[1], lines[2])
	case len(lines[0]) < 40:
		return parseTD2(lines[0], lines[1])
	case len(lines[0]) > 40:
		return parseTD3(lines[0], lines[1])
	default:
		return MRZData{}, errors.New("Unknown MRZ format")
	}
}

// Helper function to parse TD1 format
func parseTD1(line1, line2, line3 string) (MRZData, error) {
	fmt.Println("I am TD1")
	// Example fields for TD1 format
	documentType := line1[0:2]
	countryCode := line1[2:5]
	documentNumber := line1[5:14]
	nationality := line2[15:18]
	sex := string(line2[7])
	names := line3
	_ = nationality

	// Create MRZData struct
	mrzData := MRZData{
		DocumentType:   documentType,
		CountryCode:    countryCode,
		Names:          names,
		DocumentNumber: documentNumber,
		Sex:            sex,
		// Add more fields as needed
	}

	return mrzData, nil
}

// Helper function to parse TD2 format
func parseTD2(line1, line2 string) (MRZData, error) {
	fmt.Println("I am TD2")
	// Example fields for TD2 format
	documentType := line1[0:2]
	countryCode := line1[2:5]
	names := line1[5:36]
	documentNumber := line2[0:9]
	sex := string(line2[20])

	// Create MRZData struct
	mrzData := MRZData{
		DocumentType:   documentType,
		CountryCode:    countryCode,
		Names:          names,
		DocumentNumber: documentNumber,
		Sex:            sex,
		// Add more fields as needed
	}

	return mrzData, nil
}

// Helper function to parse TD3 format
func parseTD3(line1, line2 string) (MRZData, error) {
	fmt.Println("I am TD3")
	// Example fields for TD3 format
	documentType := line1[0:2]
	countryCode := line1[2:5]
	nationality := line2[9:12]
	names := line1[4:]
	documentNumber := line2[0:9]
	sex := string(line2[20])
	// dates
	birthDate := line2[13:19]

	// Create MRZData struct
	mrzData := MRZData{
		DocumentType:   documentType,
		CountryCode:    countryCode,
		Names:          names,
		DocumentNumber: documentNumber,
		Sex:            sex,
		Nationality:    nationality,
		BirthDate:      birthDate,
		// Add more fields as needed
	}

	return mrzData, nil
}
