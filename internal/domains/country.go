package domains

import "github.com/google/uuid"

type Country struct {
	ID   string
	Name CountryName
}

func NewCountry(countryName CountryName) *Country {
	return &Country{
		ID:   uuid.New().String(),
		Name: countryName,
	}
}

type CountryName string

const (
	SWEDEN = "SWEDEN"
	SPAIN  = "SPAIN"
)
