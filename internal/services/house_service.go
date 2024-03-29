package services

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/models"
	"gorm.io/gorm"
)

type HouseServiceI interface {
	Create(house *models.House) error
	FindAll() ([]models.House, error)
	FindById(id string) (models.House, error)
	SearchHouses(country string, textAreaSearchValue string, houseTypes []string, sort string) ([]models.House, error)
}

type HouseService struct{ db *gorm.DB }

func UseHouseService() HouseServiceI { return &HouseService{db: configs.GetPostgresDB()} }

func (s *HouseService) Create(house *models.House) error {
	result := s.db.Create(house)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *HouseService) FindAll() ([]models.House, error) {
	var houses []models.House
	result := s.db.Preload("HouseImages").Preload("Broker").Find(&houses)
	if result.Error != nil {
		return nil, result.Error
	}
	return houses, nil
}

func (s *HouseService) FindById(id string) (models.House, error) {
	var house models.House
	result := s.db.Preload("HouseImages").Preload("Broker").First(&house, &id)
	if result.Error != nil {
		return models.House{}, result.Error
	}
	return house, nil
}

func (s *HouseService) SearchHouses(country string, textAreaSearchValue string, houseTypes []string, sort string) ([]models.House, error) {
	var houses []models.House
	tx := s.db

	if country != "ANY" {
		tx = tx.Where("country = ?", country)
	}

	if textAreaSearchValue != "" {
		tx = tx.Where("city LIKE ?", "%"+textAreaSearchValue+"%")
	}

	if len(houseTypes) > 0 {
		tx = tx.Where("house_types IN ?", houseTypes)
	}

	switch sort {
	case "price-desc":
		tx = tx.Order("price DESC")
	case "price-asc":
		tx = tx.Order("price ASC")
	case "featured":
		tx = tx.Order("price")
	default:
		tx = tx.Order("price")
	}

	tx.Find(&houses)

	return houses, nil
}
