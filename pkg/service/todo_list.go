package service

import (
	"github.com/kalimoldayev02/api-golang/models"
	"github.com/kalimoldayev02/api-golang/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetTodoLists(userId int) ([]models.TodoList, error) {
	return s.repo.GetTodoLists(userId)
}

func (s *TodoListService) GetTodoListById(userId int, id int) (models.TodoList, error) {
	return s.repo.GetTodoListById(userId, id)
}

func (s *TodoListService) DeleteTodoList(userId int, id int) error {
	return s.repo.DeleteTodoList(userId, id)
}

func (s *TodoListService) Update(userId int, id int, input models.UpdateTodoList) error {
	if err := models.Validate(input); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
