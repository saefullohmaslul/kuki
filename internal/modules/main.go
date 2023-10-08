package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	internalGrpc "github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/modules/todos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	grpcPort := os.Getenv("GRPC_PORT")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		panic(err)
	}

	var opt []grpc.ServerOption
	server := grpc.NewServer(opt...)
	reflection.Register(server)

	todosGrpcHandler := todos.NewTodosGrpcHandler()

	internalGrpc.RegisterTodosHandlerServer(server, todosGrpcHandler)

	fmt.Printf("⇨ grpc server started on %s\n", color.New().Magenta(fmt.Sprintf(":%s", grpcPort)))

	go func() {
		err := server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:3001",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	gwmux := runtime.NewServeMux()

	err = internalGrpc.RegisterTodosHandlerHandler(context.Background(), gwmux, conn)

	if err != nil {
		panic(err)
	}

	app.Any("*", echo.WrapHandler(gwmux))

	fmt.Printf("⇨ http server started on %s\n", color.New().Magenta(":8080"))

	err = app.Start(":8080")
	if err != nil {
		panic(err)
	}
}
