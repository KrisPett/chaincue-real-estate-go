package domains

import "github.com/google/uuid"

type Broker struct {
	ID          string
	Name        string
	PhoneNumber string
	Email       string
}

func NewBroker(email string) *Broker {
	return &Broker{
		ID:          uuid.New().String(),
		Name:        "",
		PhoneNumber: "",
		Email:       email,
	}
}
