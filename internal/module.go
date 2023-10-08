package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	internalGrpc "github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/modules/todos"
	"github.com/saefullohmaslul/kuki/internal/pkg/env"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var Module = fx.Options(
	todos.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	todos todos.GrpcHandler,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				env.LoadDefaultEnv()
				grpcPort := os.Getenv("GRPC_PORT")
				grpcListen, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
				restPort := os.Getenv("REST_PORT")

				go func() {
					if err != nil {
						panic(err)
					}

					var opt []grpc.ServerOption
					server := grpc.NewServer(opt...)
					reflection.Register(server)

					internalGrpc.RegisterTodosHandlerServer(server, todos)
					fmt.Printf("⇨ grpc server started on %s\n", color.New().Magenta(fmt.Sprintf(":%s", grpcPort)))

					err = server.Serve(grpcListen)
					if err != nil {
						panic(err)
					}
				}()

				go func() {
					conn, err := grpc.DialContext(
						ctx,
						grpcListen.Addr().String(),
						grpc.WithBlock(),
						grpc.WithTransportCredentials(insecure.NewCredentials()),
					)
					if err != nil {
						panic(err)
					}

					app := echo.New()
					mux := runtime.NewServeMux()

					err = internalGrpc.RegisterTodosHandlerHandler(ctx, mux, conn)
					if err != nil {
						panic(err)
					}

					app.Any("*", echo.WrapHandler(mux))

					fmt.Printf("⇨ http server started on %s\n", color.New().Magenta(fmt.Sprintf(":%s", restPort)))

					err = app.Start(fmt.Sprintf(":%s", restPort))
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}
