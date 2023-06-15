package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kalimoldayev02/api-golang/models"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUserIdByCredentials(email, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetTodoLists(userId int) ([]models.TodoList, error)
	GetTodoListById(userId int, id int) (models.TodoList, error)
	DeleteTodoList(userId int, id int) error
	Update(userId int, id int, input models.UpdateTodoList) error
}

type TodoItem interface{}

type Repository struct {
	Auth
	TodoList
	TodoItem
}

func NewRespository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:     NewAuthRespository(db),
		TodoList: NewTodoListRepository(db),
	}
}
