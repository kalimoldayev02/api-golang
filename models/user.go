package models

type User struct {
	Id        int    `json:"-" db:"id"`
	FirstName string `json:"first_name" binding:"required"` // у gin есть свойство binding -> required
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
