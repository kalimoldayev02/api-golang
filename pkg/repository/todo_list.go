package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kalimoldayev02/api-golang/models"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) Create(userId int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin() // для транзакции
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf(
		"insert into %s (title, description) values ($1, $2) returning id",
		todoListsTable,
	)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf(
		"insert into %s (user_is, list_id) values ($1, $2)",
		usersListsTable,
	)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
