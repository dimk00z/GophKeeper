package entity

type Login struct {
	Name     string `json:"name"`
	UserName string `json:"user_hame"`
	Password string `json:"password"`
	URI      string `json:"uri"`
}
