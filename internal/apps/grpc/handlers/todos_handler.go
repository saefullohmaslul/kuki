package handlers

import (
	"context"
)

type todosGrpcHandler struct {
}

func NewTodosGrpcHandler() TodosGrpcHandler {
	return &todosGrpcHandler{}
}

// GetTodo is a function to get to do by id
func (h *todosGrpcHandler) GetTodo(ctx context.Context, params *GetTodoRequest) (data *Todo, err error) {
	data = &Todo{
		TodoId:      "abc",
		Title:       "Makan siang",
		Description: "Makan siang dengan istri",
		Completed:   false,
	}
	return
}

func (h *todosGrpcHandler) mustEmbedUnimplementedTodosHandlerServer() {
	panic("implement me")
}

func (h *todosGrpcHandler) CreateTodo(ctx context.Context, params *CreateTodoRequest) (data *CreateTodoResponse, err error) {
	return
}
func (h *todosGrpcHandler) UpdateTodo(ctx context.Context, params *UpdateTodoRequest) (data *UpdateTodoResponse, err error) {
	return
}
func (h *todosGrpcHandler) DeleteTodo(ctx context.Context, params *DeleteTodoRequest) (data *Empty, err error) {
	return
}
