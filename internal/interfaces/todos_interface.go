package interfaces

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/dtos"
	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type TodosGrpcHandler interface {
	grpc.TodosHandlerServer
}

type TodosRepository interface {
	FindTodoById(id string) (*models.Todos, error)
	UpdateTodoById(id string, request *models.Todos) error
	DeleteTodoById(id string) error
	GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error)
	CreateTodo(ctx context.Context, params *models.Todos) (data dtos.CreateTodoResponse, err error)
}

type TodosUseCase interface {
	FindTodoById(id string) (*models.Todos, error)
	UpdateTodoById(id string, request *models.Todos) error
	DeleteTodoById(id string) error
	GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error)
	CreateTodo(ctx context.Context, params *dtos.CreateTodoRequest) (data dtos.CreateTodoResponse, err error)
}
