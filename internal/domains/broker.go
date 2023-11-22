package domains

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Broker struct {
	gorm.Model
	ID          string
	Name        string
	PhoneNumber string
	Email       string
	HouseID     string
}

func NewBroker(email string) *Broker {
	return &Broker{
		ID:          uuid.New().String(),
		Name:        "",
		PhoneNumber: "",
		Email:       email,
		HouseID:     "",
	}
}
