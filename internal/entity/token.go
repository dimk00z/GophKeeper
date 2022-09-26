package entity

type JWT struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenMaxAge  int
	RefreshTokenMaxAge int
	Domain             string
}
