package entity

type Card struct {
	Name            string `json:"name"`
	CardHolderName  string `json:"card_holder_hame"`
	Number          string `json:"number"`
	Brand           string `json:"brand"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	SecurityCode    int    `json:"security_code"`
}
