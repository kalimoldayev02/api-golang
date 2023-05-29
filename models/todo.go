package models

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"descrition"`
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
