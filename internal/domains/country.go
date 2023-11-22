package domains

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
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
