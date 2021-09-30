package models

type User struct {
	Id        int64  `json:"-"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Validated bool   `json:"validated"` 
}
