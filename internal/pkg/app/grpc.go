package app

import (
	"context"
	"fmt"
	"net"

	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	internalGrpc "github.com/saefullohmaslul/kuki/internal/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Grpc struct {
	port         string
	listener     net.Listener
	dependencies *Dependencies
	server       *grpc.Server
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
	a.server = grpc.NewServer(opt...)
	reflection.Register(a.server)

	internalGrpc.RegisterTodosHandlerServer(a.server, a.dependencies.Todos)

	fmt.Printf("⇨ grpc server started on %s\n", color.New().Magenta(fmt.Sprintf(":%s", a.port)))

	err := a.server.Serve(a.listener)
	if err != nil {
		panic(err)
	}
}

func (a *Grpc) Shutdown(ctx context.Context) {
	a.server.GracefulStop()

	log.Info("gRPC server shutdown")
}
