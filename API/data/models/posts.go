package models

import "time"

type Post struct {
	Id          int64  `json:"id" validate:"required,number"`
	Time        time.Time `json:"time" validate:"required"`
	OwnerId     int64  `json:"-" validate:"required,number"`
	MediaName   string `json:"media_name"`
	MediaUrl    string `json:"media_url"`
	Description string `json:"description"`
	Likes       int64  `json:"likes" validate:"number"`
	Hearts      int64  `json:"hearts" validate:"number"`
	Laughs      int64  `json:"laughs" validate:"number"`
	Sads        int64  `json:"sads" validate:"number"`
	Angries     int64  `json:"angries" validate:"number"`
}

type PostComments struct {
	Id            int64  `json:"id" validate:"required,number"`
	PostId        int64  `json:"post_id" validate:"required,number"`
	OwnerId       int64  `json:"-" validate:"required,number"`
	OwnerUsername string `json:"-" validate:"required"`
	Comment       string `json:"comment" validate:"required"`
}
