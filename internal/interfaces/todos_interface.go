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
	DeleteTodoById(id string) error
	GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error)
	CreateTodo(ctx context.Context, params *models.Todos) (data dtos.CreateTodoResponse, err error)
	UpdateTodo(ctx context.Context, params *dtos.UpdateTodoRequest) (data dtos.UpdateTodoResponse, err error)
}

type TodosUseCase interface {
	DeleteTodoById(id string) error
	GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error)
	CreateTodo(ctx context.Context, params *dtos.CreateTodoRequest) (data dtos.CreateTodoResponse, err error)
	UpdateTodo(ctx context.Context, params *dtos.UpdateTodoRequest) (data dtos.UpdateTodoResponse, err error)
}
