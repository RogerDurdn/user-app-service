package model

type User struct {
	Id       int    `json:"id" required:"true"`
	Name     string `json:"name" required:"true"`
	Password string `json:"password" required:"true"`
}
