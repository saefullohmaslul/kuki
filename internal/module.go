package main

import (
	"context"
	"fmt"
	"github.com/saefullohmaslul/kuki/internal/modules/todos"
	"github.com/saefullohmaslul/kuki/internal/pkg/app"
	"github.com/saefullohmaslul/kuki/internal/pkg/env"
	"go.uber.org/fx"
	"net"
	"os"
)

var Module = fx.Options(
	todos.Module,
	app.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	dependencies *app.Dependencies,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				env.LoadDefaultEnv()

				var (
					grpcPort      = os.Getenv("GRPC_PORT")
					restPort      = os.Getenv("REST_PORT")
					grpcListen, _ = net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
				)

				go func() {
					grpc := app.NewGrpc(grpcPort, grpcListen, dependencies)
					grpc.Start(ctx)
				}()

				go func() {
					rest := app.NewRest(restPort, grpcListen)
					rest.Start(ctx)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}
