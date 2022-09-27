package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Name            string `gorm:"size:100"`
	CardHolderName  string `gorm:"size:100"`
	Number          string `gorm:"size:50"`
	Brand           string `gorm:"size:20"`
	ExpirationMonth int
	ExpirationYear  int
	SecurityCode    int
	UserID          uuid.UUID
}
