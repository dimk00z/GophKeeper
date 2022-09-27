package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SecretNote struct {
	gorm.Model
	Name   string `gorm:"size:100"`
	Note   string
	UserID uuid.UUID
}
