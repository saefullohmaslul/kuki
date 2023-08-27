package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/color"
	"github.com/saefullohmaslul/kuki/internal/modules/todos"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("GRPC_PORT")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	todosGrpcHandler := todos.NewTodosGrpcHandler()

	opt := []grpc.ServerOption{}

	server := grpc.NewServer(opt...)

	todos.RegisterTodosHandlerServer(server, todosGrpcHandler)

	fmt.Printf("⇨ grpc server started on %s\n", color.New().Magenta(listen.Addr()))

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
