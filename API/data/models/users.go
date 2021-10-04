package models

type User struct {
	Id        int64  `json:"-" validate:"required,number"`
	Name      string `json:"name" validate:"required"`
	Lastname  string `json:"lastname"`
	Fullname  string `json:"fullname"`
	Username  string `json:"-"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"-" validate:"required"`
	AvatarId  int64  `json:"-" validate:"required,number"`
	Validated bool   `json:"-"`
}

type UserAvatar struct {
	Id       int64  `json:"id" validate:"required,number"`
	OwnerId  int64  `json:"-" validate:"required,number"`
	FileName string `json:"file_name" validate:"required"`
}

type AuthToken struct {
	Email       string `json:"email" validate:"required,email"`
	AccessToken string `json:"access_token" validate:"required"`
	TokenType   string `json:"token_type" validate:"required"`
}
