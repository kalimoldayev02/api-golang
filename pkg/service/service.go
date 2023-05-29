package service

import "github.com/kalimoldayev02/api-golang/pkg/repository"

type Auth interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Auth // хранилище для сервисов
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service { // конструктор
	return &Service{}
}
