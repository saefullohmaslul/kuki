package interfaces

import (
	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type TodosGrpcHandler interface {
	grpc.TodosHandlerServer
}

type TodosRepository interface {
	InsertTodo(request *models.Todo) error
	FindTodoById(id string) (*models.Todo, error)
	UpdateTodoById(id string, request *models.Todo) error
	DeleteTodoById(id string) error
}

type TodosService interface {
	InsertTodo(request *models.Todo) error
	FindTodoById(id string) (*models.Todo, error)
	UpdateTodoById(id string, request *models.Todo) error
	DeleteTodoById(id string) error
}
