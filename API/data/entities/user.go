package entities

type User struct {
	Id 		 int64  `json:"-"`
	Name 	 string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"-"`
}