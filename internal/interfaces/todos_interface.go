package interfaces

import (
	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type TodosGrpcHandler interface {
	grpc.TodosHandlerServer
}

type TodosRepository interface {
	InsertTodo(request *models.Todos) error
	FindTodoById(id string) (*models.Todos, error)
	UpdateTodoById(id string, request *models.Todos) error
	DeleteTodoById(id string) error
}

type TodosUseCase interface {
	InsertTodo(request *models.Todos) error
	FindTodoById(id string) (*models.Todos, error)
	UpdateTodoById(id string, request *models.Todos) error
	DeleteTodoById(id string) error
}
