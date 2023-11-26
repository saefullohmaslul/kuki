package todos

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewUseCase),
	fx.Provide(NewGrpcHandler),
	fx.Provide(NewRestHandler),
)
