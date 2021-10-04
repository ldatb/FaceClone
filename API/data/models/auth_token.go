package models

type AuthToken struct {
	Email       string `json:"email" validate:"required,email"`
	AccessToken string `json:"access_token" validate:"required"`
	TokenType   string `json:"token_type" validate:"required"`
}
