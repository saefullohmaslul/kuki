package todos

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/grpc"
)

type grpcHandler struct {
	usecase UseCase
}

func NewGrpcHandler(usecase UseCase) GrpcHandler {
	return &grpcHandler{
		usecase: usecase,
	}
}

// GetTodo is a function to get to do by id
func (h *grpcHandler) GetTodo(ctx context.Context, params *grpc.GetTodoRequest) (data *grpc.Todo, err error) {
	data, err = h.usecase.GetTodo(ctx, params)
	return
}

func (h *grpcHandler) CreateTodo(ctx context.Context, params *grpc.CreateTodoRequest) (data *grpc.CreateTodoResponse, err error) {
	mockData := &grpc.Todo{
		Title:       params.Title,
		Description: params.Description,
	}

	response := &grpc.CreateTodoResponse{
		Todo: mockData,
	}

	return response, nil
}
func (h *grpcHandler) UpdateTodo(ctx context.Context, params *grpc.UpdateTodoRequest) (data *grpc.UpdateTodoResponse, err error) {
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
func (h *grpcHandler) DeleteTodo(ctx context.Context, params *grpc.DeleteTodoRequest) (data *grpc.Empty, err error) {
	return &grpc.Empty{}, nil
}
