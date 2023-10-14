package todos

import (
	"context"

	"github.com/google/uuid"
	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type grpcHandler struct {
	todosUseCase interfaces.TodosUseCase
}

func NewGrpcHandler(todosUseCase interfaces.TodosUseCase) interfaces.TodosGrpcHandler {
	return &grpcHandler{
		todosUseCase: todosUseCase,
	}
}

// GetTodo is a function to get to do by id
func (h *grpcHandler) GetTodo(ctx context.Context, params *grpc.GetTodoRequest) (data *grpc.Todo, err error) {
	todo, err := h.todosUseCase.FindTodoById(params.TodoId)
	if err != nil {
		return nil, err
	}

	mockData := &grpc.Todo{
		TodoId:      todo.TodoID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	return mockData, nil
}

func (h *grpcHandler) CreateTodo(ctx context.Context, params *grpc.CreateTodoRequest) (data *grpc.CreateTodoResponse, err error) {
	request := &models.Todos{
		TodoID:      uuid.New().String(),
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todosUseCase.InsertTodo(request)
	if err != nil {
		return nil, err
	}

	mockData := &grpc.Todo{
		TodoId:      request.TodoID,
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
	request := &models.Todos{
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todosUseCase.UpdateTodoById(params.TodoId, request)
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
	err = h.todosUseCase.DeleteTodoById(params.TodoId)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
