package todos

import (
	"context"
<<<<<<< HEAD
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
=======
	"github.com/saefullohmaslul/kuki/internal/dtos"

>>>>>>> main
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
<<<<<<< HEAD
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.GetTodo")
	defer trace.Finish()
	todo, err := h.todosUseCase.FindTodoById(ctx, params.TodoId)
	if err != nil {
		return nil, err
	}
=======
	todo, err := h.todosUseCase.GetTodo(ctx, &dtos.GetTodoRequest{
		TodoID: params.TodoId,
	})
>>>>>>> main

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
<<<<<<< HEAD
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.CreateTodo")
	defer trace.Finish()
	request := &models.Todos{
		TodoID:      uuid.New().String(),
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todosUseCase.InsertTodo(ctx, request)
=======
	todo, err := h.todosUseCase.CreateTodo(ctx, &dtos.CreateTodoRequest{
		Todos: models.Todos{
			Title:       params.Title,
			Description: params.Description,
			Completed:   params.Completed,
		},
	})

>>>>>>> main
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

// UpdateTodo is function to update todo by id
func (h *grpcHandler) UpdateTodo(ctx context.Context, params *grpc.UpdateTodoRequest) (data *grpc.UpdateTodoResponse, err error) {
<<<<<<< HEAD
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.UpdateTodo")
	time.Sleep(2 / time.Second)
	defer trace.Finish()
	request := &models.Todos{
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todosUseCase.UpdateTodoById(ctx, params.TodoId, request)
=======
	todo, err := h.todosUseCase.UpdateTodo(ctx, &dtos.UpdateTodoRequest{
		Todos: models.Todos{
			Title:       params.Title,
			Description: params.Description,
			Completed:   params.Completed,
		},
	})

>>>>>>> main
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
<<<<<<< HEAD
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.DeleteTodo")
	defer trace.Finish()
	err = h.todosUseCase.DeleteTodoById(ctx, params.TodoId)
	if err != nil {
		return nil, err
	}
=======
	err = h.todosUseCase.DeleteTodo(ctx, &dtos.DeleteTodoRequest{
		TodoID: params.TodoId,
	})
>>>>>>> main

	return &grpc.Empty{}, err
}
