package models

type Media struct {
	Id       int64  `json:"id" validate:"required,number"`
	OwnerId  int64  `json:"-" validate:"required,number"`
	FileName string `json:"file_name" validate:"required"`
}
