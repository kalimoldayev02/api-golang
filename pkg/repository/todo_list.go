package repository

import (
	"fmt"
	"strings"

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
		"insert into %s (user_id, list_id) values ($1, $2)",
		usersListsTable,
	)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListRepository) GetTodoLists(userId int) ([]models.TodoList, error) {
	var todoLists []models.TodoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id = $1", todoListsTable, usersListsTable)
	err := r.db.Select(&todoLists, query, userId)

	return todoLists, err
}

func (r *TodoListRepository) GetTodoListById(userId int, id int) (models.TodoList, error) {
	var todoList models.TodoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id = $1 and ul.list_id = $2", todoListsTable, usersListsTable)
	err := r.db.Get(&todoList, query, userId, id)

	return todoList, err
}

func (r *TodoListRepository) DeleteTodoList(userId int, id int) error {
	query := fmt.Sprintf("delete from %s tl using %s ul where tl.id = ul.list_id and ul.user_id = $1 and ul.list_id = $2", todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, id)

	return err
}

func (r *TodoListRepository) Update(userId int, id int, input models.UpdateTodoList) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("update %s tl set %s from %s ul where tl.id = ul.list_id and ul.list_id = %d and ul.user_id = $%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, id, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
