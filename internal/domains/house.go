package domains

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
	NumberRooms int
	Beds        int
	Price       string
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
		NumberRooms: 0,
		Beds:        0,
		Price:       "",
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
