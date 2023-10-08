package todos

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/grpc"
)

type todosGrpcHandler struct {
}

func NewTodosGrpcHandler() GrpcHandler {
	return &todosGrpcHandler{}
}

// GetTodo is a function to get to do by id
func (h *todosGrpcHandler) GetTodo(ctx context.Context, params *grpc.GetTodoRequest) (data *grpc.Todo, err error) {
	data = &grpc.Todo{
		TodoId:      "abc",
		Title:       "Makan siang",
		Description: "Makan siang dengan istri",
		Completed:   false,
	}
	return
}

func (h *todosGrpcHandler) CreateTodo(ctx context.Context, params *grpc.CreateTodoRequest) (data *grpc.CreateTodoResponse, err error) {
	mockData := &grpc.Todo{
		Title:       params.Title,
		Description: params.Description,
	}

	response := &grpc.CreateTodoResponse{
		Todo: mockData,
	}

	return response, nil
}
func (h *todosGrpcHandler) UpdateTodo(ctx context.Context, params *grpc.UpdateTodoRequest) (data *grpc.UpdateTodoResponse, err error) {
	mockData := &grpc.Todo{
		TodoId:      params.TodoId,
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	response := &grpc.UpdateTodoResponse{
		Todo: mockData,
	}

	return response, nil
}
func (h *todosGrpcHandler) DeleteTodo(ctx context.Context, params *grpc.DeleteTodoRequest) (data *grpc.Empty, err error) {
	return &grpc.Empty{}, nil
}
