package todos

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/grpc"
)

type GrpcHandler interface {
	grpc.TodosHandlerServer
}

type UseCase interface {
	GetTodo(ctx context.Context, params *grpc.GetTodoRequest) (data *grpc.Todo, err error)
}
