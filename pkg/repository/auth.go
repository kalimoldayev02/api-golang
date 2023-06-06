package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kalimoldayev02/api-golang/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRespository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"insert into %s (first_name, last_name, email, password) values ($1, $2, $3, $4) returning id",
		usersTable,
	)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUserIdByCredentials(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id from %s where email=$1 and password=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
