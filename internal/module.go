package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/saefullohmaslul/kuki/internal/modules/todos"
	"github.com/saefullohmaslul/kuki/internal/pkg/app"
	"github.com/saefullohmaslul/kuki/internal/pkg/database"
	"github.com/saefullohmaslul/kuki/internal/pkg/env"
	"github.com/saefullohmaslul/kuki/internal/pkg/jeager"
	"github.com/saefullohmaslul/kuki/internal/pkg/validator"
	"go.uber.org/fx"
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
) {
	env.LoadDefaultEnv()

	closer, err := jeager.InitializeJaeger("todos-service")
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}

	var (
		grpcPort      = os.Getenv("GRPC_PORT")
		restPort      = os.Getenv("REST_PORT")
		grpcListen, _ = net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		grpc          = app.NewGrpc(grpcPort, grpcListen, dependencies)
		rest          = app.NewRest(restPort, grpcListen)
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

				defer func() {
					if err := closer.Close(); err != nil {
						log.Printf("Could not close jaeger tracer: %s", err.Error())
					}
				}()

				rest.Shutdown(ctx)
				grpc.Shutdown(ctx)
				return nil
			},
		},
	)
}
