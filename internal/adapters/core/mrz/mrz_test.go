package mrz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentType(t *testing.T) {
	assert.Equal(t, "TD1", TD1.String())
	assert.Equal(t, "TD2", TD2.String())
	assert.Equal(t, "TD3", TD3.String())
}

func TestParseTD3(t *testing.T) {
	firstCard := `P<NGADEBOYEWA<<ADEBOYE<USMAN<<<<<<<<<<<<<<
A84824929NGA37443234M1997831<<<<<<<<<<<<<<<02`
	data, err := ParseMRZ(firstCard)
	assert.NoError(t, err)
	fmt.Printf("data: %+v\n", data)

	secondCard := `P<D<<CHEBAK<<ABOULKACIME<<<<<<<<<<<<<<<<<<<<
C47NTZ8CWOD<<8103103M2905198<<<<<<<<<<<<<<<6`
	data, err = ParseMRZ(secondCard)
	assert.NoError(t, err)
	fmt.Printf("data: %+v\n", data)
}
