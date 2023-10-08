package app

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/modules/todos"
	"go.uber.org/fx"
)

type App interface {
	Start(ctx context.Context)
}

type Dependencies struct {
	Todos todos.GrpcHandler
}

func NewDependencies(todos todos.GrpcHandler) *Dependencies {
	return &Dependencies{
		Todos: todos,
	}
}

var Module = fx.Option(
	fx.Provide(NewDependencies),
)
