package errs

import (
	"encoding/json"
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
