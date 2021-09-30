package models

type AuthAccess struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
