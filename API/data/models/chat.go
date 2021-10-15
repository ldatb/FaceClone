package models

import "time"

type Chat struct {
	Id    int64 `json:"chat-id" validate:"required,number"`
	User1 int64 `json:"-" validate:"required,number"`
	User2 int64 `json:"-" validate:"required,number"`
}

type ChatMessages struct {
	Id     int64     `json:"message-id" validate:"required,number"`
	ChatId int64     `json:"chat-id" validate:"required,number"`
	Time   time.Time `json:"time" validate:"required,number"`
	UserId int64     `json:"-" validate:"required,number"`
	Content string `json:"content" validate:"required"`
}
