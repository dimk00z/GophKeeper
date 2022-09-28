package entity

import "github.com/google/uuid"

type Login struct {
	ID       uuid.UUID
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	URI      string `json:"uri"`
}
