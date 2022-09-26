package errs

import (
	"encoding/json"
	"errors"
)

var (
	ErrWrongEmail         = errors.New("incorrect email given")
	ErrEmailAlreadyExists = errors.New("given email already exists")
	ErrWrongCredentials   = errors.New("wrong credentials have been given")
	ErrTokenValidation    = errors.New("token validation error")
	ErrUnexpectedError    = errors.New("some unexpected error")
)

type GormErr struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func ParsePostgresErr(dbErr error) (newError GormErr) {
	byteErr, err := json.Marshal(dbErr)
	if err != nil {
		return
	}

	if err = json.Unmarshal((byteErr), &newError); err != nil {
		return GormErr{}
	}

	return
}
