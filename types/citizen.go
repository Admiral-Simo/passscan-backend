package types

// Document represents parsed MRZ data
type Document struct {
	DocumentNumber string `json:"documentNumber" gorm:"column:document_number;primaryKey;uniqueIndex"`
	DocumentType   string `json:"documentType" gorm:"column:document_type"`
	CountryCode    string `json:"countryCode" gorm:"column:country_code"`
	FirstName      string `json:"firstName" gorm:"column:first_name"`
	LastName       string `json:"lastName" gorm:"column:last_name"`
	Sex            string `json:"sex" gorm:"column:sex"`
	BirthDate      string `json:"birthDate" gorm:"column:birth_date"`
	ExpireDate     string `json:"expireDate" gorm:"column:expire_date"`
}
