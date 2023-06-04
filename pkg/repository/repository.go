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
}

type TodoItem interface {
}

type Repository struct {
	Auth
	TodoList
	TodoItem
}

func NewRespository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAutRespository(db),
	}
}
