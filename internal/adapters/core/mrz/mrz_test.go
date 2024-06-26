package mrz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJapanese(t *testing.T) {
	// japanese passport
	thirdCard := `P<JPNGAIMU<<HANAKO<<<<<<<<<<<<<<<<<<<<<<<<<<
TT15442512JPN9905050F3012257<<<<<<<<<<<<<<06`
	data, err := ParseMRZ(thirdCard)
	assert.NoError(t, err)
	assert.Equal(t, "HANAKO", data.FirstName)
	assert.Equal(t, "GAIMU", data.LastName)
	assert.Equal(t, "F", data.Sex)
	assert.Equal(t, "JPN", data.CountryCode)
	assert.Equal(t, "P", data.DocumentType)
}

func TestNigeria(t *testing.T) {
	firstCard := `P<NGADEBOYEWA<<ADEBOYE<USMAN<<<<<<<<<<<<<<
A84824929NGA37443234M1997831<<<<<<<<<<<<<<<02`
	data, err := ParseMRZ(firstCard)
	assert.NoError(t, err)
	assert.Equal(t, "ADEBOYE USMAN", data.FirstName)
	assert.Equal(t, "DEBOYEWA", data.LastName)
	assert.Equal(t, "M", data.Sex)
	assert.Equal(t, "NGA", data.CountryCode)
	assert.Equal(t, "P", data.DocumentType)
}

func TextDutch(t *testing.T) {
	secondCard := `P<D<<CHEBAK<<ABOULKACIME<<<<<<<<<<<<<<<<<<<<
C47NTZ8CWOD<<8103103M2905198<<<<<<<<<<<<<<<6`
	data, err := ParseMRZ(secondCard)
	assert.NoError(t, err)
	assert.Equal(t, "ABOULKACIME", data.FirstName)
	assert.Equal(t, "CHEBAK", data.LastName)
	assert.Equal(t, "M", data.Sex)
	assert.Equal(t, "D", data.CountryCode)
	assert.Equal(t, "P", data.DocumentType)
}

// Carte National (ID Cards)

func TestIDCards(t *testing.T) {
	firstCard := `I<MARVSI99999<8K01234567
7811296M2909093MAR<<<<<<<<<<<4
TEMSAMANI<<MOUHCINE<<<<<<<<<<<<`
	data, err := ParseMRZ(firstCard)
	assert.NoError(t, err)
	fmt.Printf("data: %+v\n", data)

	secondCard := `I<MARVSI99998<5U1234567<<<<<<<
8312055F2907228MAR<<<<<<<<<<8
EL<ALAMI<ZAINED<<<<<<<<<<<<<<<<`
	data, err = ParseMRZ(secondCard)
	assert.NoError(t, err)
	fmt.Printf("data: %+v\n", data)

	simoCard := `IDMAROPI4JV82<9I776494<<<<<<<<
0604109M3202229MAR
KHALIS<<MOHAMED<<<<<<<<<<<<<<`
	data, err = ParseMRZ(simoCard)
	assert.NoError(t, err)
	fmt.Printf("data: %+v\n", data)

	fatimaCard := `IDMAROPI4JV5D<1I538650<<<<<<<<
6906148F3202229MAR<<<<<<<<<<<<9
OUKBICH<<FATIMA<<<<<<<<<<<<<<`
	data, err = ParseMRZ(fatimaCard)
	assert.NoError(t, err)
	fmt.Printf("data: %+v\n", data)
}
