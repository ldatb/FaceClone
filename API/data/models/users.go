package models

type User struct {
	Id        int64  `json:"-" validate:"required,number"`
	Name      string `json:"-" validate:"required"`
	Lastname  string `json:"-"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"-" validate:"required"`
	AvatarId  int64  `json:"-" validate:"required,number"`
	Validated bool   `json:"-"`
}

type UserAvatar struct {
	OwnerId  int64  `json:"-" validate:"required,number"`
	Id       int64  `json:"avatar_id" validate:"required,number"`
	FileName string `json:"file_name" validate:"required"`
}

type AuthToken struct {
	Email       string `json:"email" validate:"required,email"`
	AccessToken string `json:"access_token" validate:"required"`
	TokenType   string `json:"token_type" validate:"required"`
}

type UserReactedPosts struct {
	OwnerId  int64  `json:"-" validate:"required,number"`
	PostId   int64  `json:"post_id" validate:"required,number"`
	Reaction string `json:"reaction" validate:"required"`
}
