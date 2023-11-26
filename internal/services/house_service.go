package services

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/models"
	"gorm.io/gorm"
)

type HouseServiceI interface {
	Create(houseTypes models.HouseTypes, src string) error
	FindAll() ([]models.House, error)
	FindById(id string) (models.House, error)
}

type HouseService struct{ db *gorm.DB }

func UseHouseService() HouseServiceI { return &HouseService{db: configs.GetDB()} }

func (s *HouseService) Create(houseTypes models.HouseTypes, src string) error {
	house := models.NewHouse(houseTypes, src)
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
