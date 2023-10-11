package todos

import (
	"github.com/saefullohmaslul/kuki/internal/pkg/database"
	"go.uber.org/fx"
)


var Module = fx.Options(
	fx.Provide(database.NewDatabase),
	fx.Provide(NewTodosRepository),
	fx.Provide(NewTodosService),
	fx.Provide(NewGrpcHandler),
)
