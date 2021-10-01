package models

type File struct {
	Id int64 `json:"id"`
	OwnerId int64 `json:"-" validate:"required"`
	OwnerEmail string `json:"owner_email" validade:"required,email"`
	FileName string `json:"file_name" validate:"required"`
}