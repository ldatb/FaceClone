package models

type UserAvatar struct {
	Id      int64 `json:"id"`
	OwnerId int64 `json:"-" validate:"required"`
	FileName string `json:"file_name" validate:"required"`
}