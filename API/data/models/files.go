package models

type File struct {
	Id       int64  `json:"id" validate:"required"`
	OwnerId  int64  `json:"-" validate:"required"`
	FileName string `json:"file_name" validate:"required"`
}
