package interfaces

import "github.com/saefullohmaslul/kuki/internal/models"

type TodosRepository interface {
	InsertTodo(request *models.Todo) (*models.Todo, error)
}

type TodosService interface {
	InsertTodo(request *models.Todo) (*models.Todo, error)
}
