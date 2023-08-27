package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/color"
	"github.com/saefullohmaslul/kuki/internal/apps/grpc/handlers"
	"google.golang.org/grpc"
	"net"
	"os"
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

	todosGrpcHandler := handlers.NewTodosGrpcHandler()

	opt := []grpc.ServerOption{}

	server := grpc.NewServer(opt...)
	server.RegisterService(&grpc.ServiceDesc{ServiceName: "kuki.KukiService", Methods: []grpc.MethodDesc{
		{MethodName: "GetTodo", Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (data interface{}, err error) {
			return
		}},
	}}, nil)

	handlers.RegisterTodosHandlerServer(server, todosGrpcHandler)

	fmt.Printf("⇨ grpc server started on %s\n", color.New().Magenta(listen.Addr()))

	err = server.Serve(listen)
	return
}
