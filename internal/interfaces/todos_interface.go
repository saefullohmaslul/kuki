package interfaces

import (
	"context"

	"github.com/saefullohmaslul/kuki/internal/dtos"
	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type (
	// TodosGrpcHandler is an interface for todos grpcHandler
	TodosGrpcHandler interface {
		grpc.TodosHandlerServer
	}

	// TodosRepository is an interface for todos repository
	TodosRepository interface {
		GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error)
		CreateTodo(ctx context.Context, params *models.Todos) (data dtos.CreateTodoResponse, err error)
		UpdateTodo(ctx context.Context, params *dtos.UpdateTodoRequest) (data dtos.UpdateTodoResponse, err error)
		DeleteTodo(ctx context.Context, params *dtos.DeleteTodoRequest) (err error)
	}

	// TodosUseCase is an interface for todos useCase
	TodosUseCase interface {
		GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error)
		CreateTodo(ctx context.Context, params *dtos.CreateTodoRequest) (data dtos.CreateTodoResponse, err error)
		UpdateTodo(ctx context.Context, params *dtos.UpdateTodoRequest) (data dtos.UpdateTodoResponse, err error)
		DeleteTodo(ctx context.Context, params *dtos.DeleteTodoRequest) (err error)
	}
)
