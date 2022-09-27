package entity

import "github.com/google/uuid"

type Login struct {
	ID       uuid.UUID
	Name     string `json:"name"`
	UserName string `json:"user_hame"`
	Password string `json:"password"`
	URI      string `json:"uri"`
}
