package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primary_key"`
	Name            string    `gorm:"size:100"`
	CardHolderName  string
	Number          string
	Brand           string
	ExpirationMonth string
	ExpirationYear  string
	SecurityCode    string
	UserID          uuid.UUID
}
