package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Name   string    `gorm:"size:100"`
	Note   string
	UserID uuid.UUID
}
