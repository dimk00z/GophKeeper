package models

import "gorm.io/gorm"

type SavedLogin struct {
	gorm.Model
	Login    string
	Password string
}
