package entity

type Card struct {
	Name            string `json:"name"`
	CardHolderName  string `json:"card_holder_hame"`
	Number          string `json:"number"`
	Brand           string `json:"brand"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	SecurityCode    string `json:"security_code"`
}
