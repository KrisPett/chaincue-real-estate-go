package services

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/models"
	"gorm.io/gorm"
)

type HouseImageServiceI interface {
	Create(houseImage *models.HouseImage) error
	FindAll() ([]models.HouseImage, error)
	FindById(id string) (models.HouseImage, error)
}

type HouseImageService struct{ db *gorm.DB }

func UseHouseImageService() HouseImageServiceI {
	return &HouseImageService{db: configs.GetPostgresDB()}
}

func (s *HouseImageService) Create(houseImage *models.HouseImage) error {
	result := s.db.Create(houseImage)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *HouseImageService) FindAll() ([]models.HouseImage, error) {
	var houseImages []models.HouseImage
	result := s.db.Find(&houseImages)
	if result.Error != nil {
		return nil, result.Error
	}
	return houseImages, nil
}

func (s *HouseImageService) FindById(id string) (models.HouseImage, error) {
	var houseImage models.HouseImage
	result := s.db.First(&houseImage, id)
	if result.Error != nil {
		return models.HouseImage{}, result.Error
	}
	return houseImage, nil
}
