package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HouseImage struct {
	gorm.Model
	ID      string
	URL     string
	HouseID string
}

func NewHouseImage(url string, houseID string) *HouseImage {
	return &HouseImage{
		ID:      uuid.New().String(),
		URL:     url,
		HouseID: houseID,
	}
}
