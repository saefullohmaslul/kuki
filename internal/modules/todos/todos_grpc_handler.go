package todos

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
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
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.GetTodo")
	defer trace.Finish()
	todo, err := h.todosUseCase.FindTodoById(ctx, params.TodoId)
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
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.CreateTodo")
	defer trace.Finish()
	request := &models.Todos{
		TodoID:      uuid.New().String(),
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todosUseCase.InsertTodo(ctx, request)
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
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.UpdateTodo")
	time.Sleep(2 / time.Second)
	defer trace.Finish()
	request := &models.Todos{
		Title:       params.Title,
		Description: params.Description,
		Completed:   params.Completed,
	}

	err = h.todosUseCase.UpdateTodoById(ctx, params.TodoId, request)
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
	trace, ctx := opentracing.StartSpanFromContext(ctx, "grpcHandler.DeleteTodo")
	defer trace.Finish()
	err = h.todosUseCase.DeleteTodoById(ctx, params.TodoId)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
