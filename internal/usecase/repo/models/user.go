package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	CreditCards []CreditCard
	SavedLogins []SavedLogin
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user *User) ToString() string {
	return fmt.Sprintf("id: %v\nemail: %s", user.ID, user.Email)
}
