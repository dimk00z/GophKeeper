package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	OwnerName string `gorm:"size:100"`
	Number    string `gorm:"size:50"`
	Type      string `gorm:"size:20"`
	UserID    uint
}
