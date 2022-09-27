package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SavedLogin struct {
	gorm.Model
	Name     string `gorm:"size:100"`
	URI      string
	Login    string
	Password string
	UserID   uuid.UUID
}
