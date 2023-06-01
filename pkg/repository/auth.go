package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kalimoldayev02/api-golang/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAutRespository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	return 0, nil
}
