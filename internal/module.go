package main

import (
	"context"
	"fmt"
	"github.com/saefullohmaslul/kuki/internal/modules/todos"
	"github.com/saefullohmaslul/kuki/internal/pkg/app"
	"github.com/saefullohmaslul/kuki/internal/pkg/database"
	"github.com/saefullohmaslul/kuki/internal/pkg/env"
	"github.com/saefullohmaslul/kuki/internal/pkg/validator"
	"go.uber.org/fx"
	"net"
	"os"
	"time"
)

var Module = fx.Options(
	database.Module,
	validator.Module,
	todos.Module,
	app.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	dependencies *app.Dependencies,
	handlers *app.Handlers,
) {
	env.LoadDefaultEnv()

	var (
		grpcPort      = os.Getenv("GRPC_PORT")
		restPort      = os.Getenv("REST_PORT")
		grpcListen, _ = net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		grpc          = app.NewGrpc(grpcPort, grpcListen, dependencies)
		rest          = app.NewRest(restPort, handlers)
	)

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					grpc.Start(ctx)
				}()

				go func() {
					rest.Start(ctx)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
				defer cancel()

				rest.Shutdown(ctx)
				grpc.Shutdown(ctx)
				return nil
			},
		},
	)
}
