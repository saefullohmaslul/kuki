package todos

import (
	"context"

	"github.com/google/uuid"
	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type grpcHandler struct {
		todoService interfaces.TodosService

}

func NewGrpcHandler(todoService interfaces.TodosService) GrpcHandler {
	return &grpcHandler{
		todoService: todoService,
	}
}

// GetTodo is a function to get to do by id
func (h *grpcHandler) GetTodo(ctx context.Context, params *grpc.GetTodoRequest) (data *grpc.Todo, err error) {
	data = &grpc.Todo{
		TodoId:      "abc",
		Title:       "Makan siang",
		Description: "Makan siang dengan istri",
		Completed:   false,
	}
	return
}

func (h *grpcHandler) CreateTodo(ctx context.Context, params *grpc.CreateTodoRequest) (data *grpc.CreateTodoResponse, err error) {
	request := &models.Todo{
		TodoId:      uuid.New().String(),
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todoService.InsertTodo(request)
	if err != nil {
		return nil, err
	}

	mockData := &grpc.Todo{
		TodoId:      request.TodoId,
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
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
