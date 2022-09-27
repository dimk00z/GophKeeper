package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Name            string `gorm:"size:100"`
	CardHolderName  string
	Number          string
	Brand           string
	ExpirationMonth string
	ExpirationYear  string
	SecurityCode    string
	UserID          uuid.UUID
}
