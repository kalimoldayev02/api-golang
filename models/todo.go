package models

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"descrition"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	Item   int
}

type UpdateTodoList struct {
	Title       *string `json:"title"` // тип указатель, в случае чего nil
	Description *string `json:"description"`
}

func Validate(i UpdateTodoList) error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
