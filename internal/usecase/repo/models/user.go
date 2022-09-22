package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	Name           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	HashedPassword string
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.NewString()

	return
}
