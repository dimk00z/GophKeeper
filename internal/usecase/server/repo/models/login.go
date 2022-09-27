package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SavedLogin struct {
	gorm.Model
	Name     string `gorm:"size:100"`
	URI      string
	Login    string `gorm:"size:100"`
	Password string `gorm:"size:100"`
	UserID   uuid.UUID
}
