package todos

import (
	"context"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
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
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.GetTodo")
	defer trace.Finish()
	if err != nil {
		return nil, err
	}
	req := &dtos.GetTodoRequest{
		TodoID: params.TodoId,
	}
	todo, err := h.todosUseCase.GetTodo(ctx, req)

	data = &grpc.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	return
}

// CreateTodo is function to create new todo
func (h *grpcHandler) CreateTodo(ctx context.Context, params *grpc.CreateTodoRequest) (data *grpc.CreateTodoResponse, err error) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.CreateTodo")
	defer trace.Finish()
	request := &dtos.CreateTodoRequest{
		Todos: models.Todos{
			TodoID:      uuid.New().String(),
			Title:       params.Title,
			Description: params.Description,
			Completed:   params.Completed,
		},
	}

	todo, err := h.todosUseCase.CreateTodo(ctx, request)
	if err != nil {
		return nil, err
	}

	data = &grpc.CreateTodoResponse{
		Todo: &grpc.Todo{
			TodoId:      todo.TodoID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		},
	}

	return data, nil
}

// UpdateTodo is function to update todo by id
func (h *grpcHandler) UpdateTodo(ctx context.Context, params *grpc.UpdateTodoRequest) (data *grpc.UpdateTodoResponse, err error) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.UpdateTodo")
	defer trace.Finish()

	request := &dtos.UpdateTodoRequest{
		Todos: models.Todos{
			Title:       params.Title,
			Description: params.Description,
			Completed:   params.Completed,
		},
	}
	todo, err := h.todosUseCase.UpdateTodo(ctx, request)

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

// DeleteTodo is function to delete todo by id
func (h *grpcHandler) DeleteTodo(ctx context.Context, params *grpc.DeleteTodoRequest) (data *grpc.Empty, err error) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.DeleteTodo")
	defer trace.Finish()
	err = h.todosUseCase.DeleteTodo(ctx, &dtos.DeleteTodoRequest{
		TodoID: params.TodoId,
	})

	return &grpc.Empty{}, err
}
