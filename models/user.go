package models

type User struct {
	Id        int    `json:"-"`
	FirstName string `json:"first_name" binding:"required"` // у gin есть свойство binding -> required
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
