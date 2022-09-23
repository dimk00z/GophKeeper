package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	CreditCards []CreditCard
	SavedLogins []SavedLogin
}

func (user *User) ToString() string {
	return fmt.Sprintf("id: %v\nemail: %s", user.ID, user.Email)
}
