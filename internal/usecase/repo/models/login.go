package models

import (
	"gorm.io/gorm"
)

type SavedLogin struct {
	gorm.Model
	UserID   uint
	Name     string `gorm:"size:100"`
	URI      string `gorm:"size:200"`
	Login    string `gorm:"size:100"`
	Password string `gorm:"size:100"`
	Note     string `gorm:"size:500"`
}
