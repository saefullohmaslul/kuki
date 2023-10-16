package app

import (
	"context"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	internalGrpc "github.com/saefullohmaslul/kuki/internal/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Rest struct {
	port     string
	listener net.Listener
	server   *echo.Echo
}

func NewRest(port string, listener net.Listener) App {
	return &Rest{
		port:     port,
		listener: listener,
	}
}

func (a *Rest) Start(ctx context.Context) {
	conn, err := grpc.DialContext(
		ctx,
		a.listener.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("Could not dial gRPC server: %s", err.Error())
		return
	}
	if err != nil {
		panic(err)
	}

	a.server = echo.New()
	mux := runtime.NewServeMux()

	err = internalGrpc.RegisterTodosHandlerHandler(ctx, mux, conn)
	if err != nil {
		panic(err)
	}

	a.server.Any("*", echo.WrapHandler(mux))

	fmt.Printf("⇨ http server started on %s\n", color.New().Magenta(fmt.Sprintf(":%s", a.port)))

	_ = a.server.Start(fmt.Sprintf(":%s", a.port))
}

func (a *Rest) Shutdown(ctx context.Context) {
	err := a.server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Info("Rest server shutdown")
}
