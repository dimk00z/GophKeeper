package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaNote struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name   string
	Value  string
	NoteID uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type Note struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name   string    `gorm:"size:100"`
	Note   string
	UserID uuid.UUID
	Meta   []MetaNote
}
