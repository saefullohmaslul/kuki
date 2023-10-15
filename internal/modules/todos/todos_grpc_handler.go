package todos

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/dtos"

	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type grpcHandler struct {
	todosUseCase interfaces.TodosUseCase
}

// NewGrpcHandler is function to create new grpc handler
func NewGrpcHandler(todosUseCase interfaces.TodosUseCase) interfaces.TodosGrpcHandler {
	return &grpcHandler{
		todosUseCase: todosUseCase,
	}
}

// GetTodo is function to get todo by id
func (h *grpcHandler) GetTodo(ctx context.Context, params *grpc.GetTodoRequest) (data *grpc.Todo, err error) {
	todo, err := h.todosUseCase.GetTodo(ctx, &dtos.GetTodoRequest{
		TodoID: params.TodoId,
	})

	data = &grpc.Todo{
		TodoId:      todo.TodoID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	return
}

// CreateTodo is function to create new todo
func (h *grpcHandler) CreateTodo(ctx context.Context, params *grpc.CreateTodoRequest) (data *grpc.CreateTodoResponse, err error) {
	todo, err := h.todosUseCase.CreateTodo(ctx, &dtos.CreateTodoRequest{
		Todos: models.Todos{
			Title:       params.Title,
			Description: params.Description,
			Completed:   params.Completed,
		},
	})

	if err != nil {
		return
	}

	data.Todo = &grpc.Todo{
		TodoId:      todo.TodoID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	return
}

func (h *grpcHandler) UpdateTodo(ctx context.Context, params *grpc.UpdateTodoRequest) (data *grpc.UpdateTodoResponse, err error) {
	todo, err := h.todosUseCase.UpdateTodo(ctx, &dtos.UpdateTodoRequest{
		Todos: models.Todos{
			Title:       params.Title,
			Description: params.Description,
			Completed:   params.Completed,
		},
	})

	if err != nil {
		return
	}

	data.Todo = &grpc.Todo{
		TodoId:      todo.TodoID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	return
}

func (h *grpcHandler) DeleteTodo(ctx context.Context, params *grpc.DeleteTodoRequest) (data *grpc.Empty, err error) {
	err = h.todosUseCase.DeleteTodoById(params.TodoId)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
