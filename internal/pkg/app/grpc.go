package app

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/color"
	internalGrpc "github.com/saefullohmaslul/kuki/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Grpc struct {
	port         string
	listener     net.Listener
	dependencies *Dependencies
}

func NewGrpc(port string, listener net.Listener, dependencies *Dependencies) App {
	return &Grpc{
		port:         port,
		listener:     listener,
		dependencies: dependencies,
	}
}

func (a *Grpc) Start(ctx context.Context) {
	var opt []grpc.ServerOption
	server := grpc.NewServer(opt...)
	reflection.Register(server)

	internalGrpc.RegisterTodosHandlerServer(server, a.dependencies.Todos)

	fmt.Printf("⇨ grpc server started on %s\n", color.New().Magenta(fmt.Sprintf(":%s", a.port)))

	err := server.Serve(a.listener)
	if err != nil {
		panic(err)
	}
}
