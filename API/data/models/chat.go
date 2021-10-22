package models

import "time"

type Chat struct {
	Id    int64 `json:"chat-id" validate:"required,number"`
	User1 int64 `json:"-" validate:"required,number"`
	User2 int64 `json:"-" validate:"required,number"`
}

type ChatMessages struct {
	Id      int64     `json:"-" validate:"required,number"`
	ChatId  int64     `json:"-" validate:"required,number"`
	Time    time.Time `json:"time" validate:"required,number"`
	UserId  int64     `json:"-" validate:"required,number"`
	Content string    `json:"content" validate:"required"`
}

type ChatImages struct {
	Id        int64     `json:"-" validate:"required,number"`
	ChatId    int64     `json:"-" validate:"required,number"`
	Time      time.Time `json:"time" validate:"required,number"`
	UserId    int64     `json:"-" validate:"required,number"`
	ImageName string    `json:"image-name" validate:"required"`
	ImageDir  string    `json:"image-dir" validate:"required"`
}
