package model

type TokenResponse struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken string `json:"access_token"`
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}