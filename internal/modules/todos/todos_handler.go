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
	todo, err := h.todoService.FindTodoById(params.TodoId)
	if err != nil {
		return nil, err
	}

	mockData := &grpc.Todo{
		TodoId:      todo.TodoId,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	return mockData, nil
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
	request := &models.Todo{
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todoService.UpdateTodoById(params.TodoId, request)
	if err != nil {
		return nil, err
	}

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
	err = h.todoService.DeleteTodoById(params.TodoId)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
