package entity

import "github.com/google/uuid"

type SecretNote struct {
	ID   uuid.UUID
	Name string `json:"name"`
	Note string `json:"password"`
}
