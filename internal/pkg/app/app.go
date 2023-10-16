package app

import (
	"context"

	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"go.uber.org/fx"
)

type App interface {
	Start(ctx context.Context)
	Shutdown(ctx context.Context)
}

type Dependencies struct {
	Todos interfaces.TodosGrpcHandler
}

func NewDependencies(todos interfaces.TodosGrpcHandler) *Dependencies {
	return &Dependencies{
		Todos: todos,
	}
}

var Module = fx.Option(
	fx.Provide(NewDependencies),
)
