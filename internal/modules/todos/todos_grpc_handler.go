package todos

import (
	"context"
)

type grpcHandler struct {
}

func NewGrpcHandler() GrpcHandler {
	return &grpcHandler{}
}

// GetTodo is a function to get to do by id
func (h *grpcHandler) GetTodo(ctx context.Context, params *GetTodoRequest) (data *Todo, err error) {
	data = &Todo{
		TodoId:      "abc",
		Title:       "Makan siang",
		Description: "Makan siang dengan istri",
		Completed:   false,
	}
	return
}

func (h *grpcHandler) mustEmbedUnimplementedTodosHandlerServer() {
	panic("implement me")
}

func (h *grpcHandler) CreateTodo(ctx context.Context, params *CreateTodoRequest) (data *CreateTodoResponse, err error) {
	return
}
func (h *grpcHandler) UpdateTodo(ctx context.Context, params *UpdateTodoRequest) (data *UpdateTodoResponse, err error) {
	return
}
func (h *grpcHandler) DeleteTodo(ctx context.Context, params *DeleteTodoRequest) (data *Empty, err error) {
	return
}
