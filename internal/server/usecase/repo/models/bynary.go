package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaBynary struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string
	Value    string
	BynaryID uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type Bynary struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name   string
	Bynary []byte
	UserID uuid.UUID
	Meta   []MetaBynary
}
