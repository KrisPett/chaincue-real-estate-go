package services

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/models"
	"gorm.io/gorm"
)

type CountryServiceI interface {
	Create(countryName models.CountryName) error
	FindAll() ([]models.Country, error)
}

type CountryService struct{ db *gorm.DB }

func UseCountryService() CountryServiceI { return &CountryService{db: configs.GetPostgresDB()} }

func (s *CountryService) Create(countryName models.CountryName) error {
	broker := models.NewCountry(countryName)
	result := s.db.Create(broker)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *CountryService) FindAll() ([]models.Country, error) {
	var countries []models.Country
	result := s.db.Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}
