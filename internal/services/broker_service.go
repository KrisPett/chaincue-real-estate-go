package services

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/models"
	"gorm.io/gorm"
)

type BrokerServiceI interface {
	Create(email string) error
}

type BrokerService struct{ db *gorm.DB }

func UseBrokerService() BrokerServiceI { return &BrokerService{db: configs.GetDB()} }

func (s *BrokerService) Create(email string) error {
	broker := models.NewBroker(email)
	result := s.db.Create(broker)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
