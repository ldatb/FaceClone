package models

type User struct {
	Id         int64  `json:"-" validate:"required,number"`
	Name       string `json:"-" validate:"required"`
	Lastname   string `json:"-"`
	Fullname   string `json:"fullname"`
	Username   string `json:"username"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"-" validate:"required"`
	AvatarFile string `json:"avatar_file" validate:"required"`
	AvatarUrl  string `json:"avatar_url" validate:"required"`
	Validated  bool   `json:"-"`
	Followers  int    `json:"followers"`
	Following  int    `json:"following"`
	Friends    int    `json:"friends"`
	AccessToken string `json:"-"`
}

type UserFriends struct {
	OwnerId   int64    `json:"-" validate:"required,number"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	Friends   []string `json:"friends"`
}

type UserAuthToken struct {
	Email       string `json:"email" validate:"required,email"`
	AuthToken string `json:"auth_token" validate:"required"`
	TokenType   string `json:"token_type" validate:"required"`
}

type UserReactedPosts struct {
	OwnerId  int64  `json:"-" validate:"required,number"`
	PostId   int64  `json:"post_id" validate:"required,number"`
	Reaction string `json:"reaction" validate:"required"`
}