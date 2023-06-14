package service

import (
	"github.com/kalimoldayev02/api-golang/models"
	"github.com/kalimoldayev02/api-golang/pkg/repository"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetTodoLists(userId int) ([]models.TodoList, error)
	GetTodoListById(userId int, id int) (models.TodoList, error)
}

type TodoItem interface{}

type Service struct {
	Auth // хранилище для сервисов
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service { // конструктор
	return &Service{
		Auth:     NewAuthService(repo.Auth),
		TodoList: NewTodoListService(repo.TodoList),
	}
}
