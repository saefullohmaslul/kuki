package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	internalGrpc "github.com/saefullohmaslul/kuki/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

type Rest struct {
	port     string
	listener net.Listener
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
		panic(err)
	}

	app := echo.New()
	mux := runtime.NewServeMux()

	err = internalGrpc.RegisterTodosHandlerHandler(ctx, mux, conn)
	if err != nil {
		panic(err)
	}

	app.Any("*", echo.WrapHandler(mux))

	fmt.Printf("⇨ http server started on %s\n", color.New().Magenta(fmt.Sprintf(":%s", a.port)))

	err = app.Start(fmt.Sprintf(":%s", a.port))
	if err != nil {
		panic(err)
	}
}
