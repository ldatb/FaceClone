package models

type User struct {
	Id        int64  `json:"-" validate:"required number"`
	Name      string `json:"name" validate:"required"`
	Lastname  string `json:"lastname"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"-" validate:"required"`
	AvatarId  int64  `json:"-" validate:"required"`
	Validated bool   `json:"-"`
}
