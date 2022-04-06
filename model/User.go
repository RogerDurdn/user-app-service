package model

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName" required:"true" gorm:"unique"`
	Password string `json:"password" required:"true"`
}
