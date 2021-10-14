package models

type UserFriends struct {
	OwnerId   int64    `json:"-" validate:"required,number"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	Friends   []string `json:"friends"`
}
