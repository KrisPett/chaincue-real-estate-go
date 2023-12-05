package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	ID          string
	Title       string
	Description string
	Location    string
	Country     string
	City        string
	NumberRooms int
	Beds        int
	Price       int
	Src         string `gorm:"not null"`
	Sold        bool
	HouseTypes  HouseTypes   `gorm:"not null"`
	HouseImages []HouseImage `gorm:"foreignKey:HouseID"`
	Broker      *Broker      `gorm:"foreignKey:HouseID"`
}

func NewHouse(houseTypes HouseTypes, src string) *House {
	return &House{
		ID:          uuid.New().String(),
		Title:       "",
		Description: "",
		Location:    "",
		Country:     "",
		City:        "",
		NumberRooms: 0,
		Beds:        0,
		Price:       0,
		Src:         src,
		Sold:        false,
		HouseTypes:  houseTypes,
		HouseImages: []HouseImage{},
		Broker:      nil,
	}
}

type HouseTypes string

const (
	CONDOMINIUM       = "CONDOMINIUM"
	VILLA             = "VILLA"
	TOWNHOUSE         = "TOWNHOUSE"
	VACATION_HOME     = "VACATION_HOME"
	ESTATES_AND_FARMS = "ESTATES_AND_FARMS"
	LAND              = "LAND"
	OTHER_HOUSES      = "OTHER_HOUSES"
)
