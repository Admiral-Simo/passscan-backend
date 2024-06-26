package types

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
}
