package repository

type Auth interface {
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

func NewRespository() *Repository {
	return &Repository{}
}
