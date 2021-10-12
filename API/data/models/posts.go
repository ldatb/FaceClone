package models

type Post struct {
	Id          int64  `json:"id" validate:"required,number"`
	OwnerId     int64  `json:"-" validate:"required,number"`
	MediaId     int64  `json:"media_id" validate:"number"`
	Description string `json:"description"`
}

type PostMedia struct {
	Id       int64  `json:"id" validate:"required,number"`
	PostId   int64  `json:"post_id" validate:"required,number"`
	OwnerId  int64  `json:"-" validate:"required,number"`
	FileName string `json:"file_name" validate:"required"`
}

type PostComments struct {
	Id            int64  `json:"id" validate:"required,number"`
	PostId        int64  `json:"post_id" validate:"required,number"`
	OwnerId       int64  `json:"-" validate:"required,number"`
	OwnerUsername string `json:"-" validate:"required"`
	Comment       string `json:"comment" validate:"required"`
}
