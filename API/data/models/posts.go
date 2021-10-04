package models

type Post struct {
	Id          int64  `json:"id" validate:"required,number"`
	OwnerId     int64  `json:"-" validate:"required,number"`
	MediaId     int64  `json:"media_id" validate:"number"`
	Description string `json:"description" validate:"alphanumunicode"`
}

type Media struct {
	Id       int64  `json:"id" validate:"required,number"`
	OwnerId  int64  `json:"-" validate:"required,number"`
	FileName string `json:"file_name" validate:"required"`
}
