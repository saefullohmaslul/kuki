package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"reflect"
)

type Rest struct {
	port     string
	handlers *Handlers
	server   *echo.Echo
}

func NewRest(port string, handlers *Handlers) App {
	return &Rest{
		port:     port,
		handlers: handlers,
	}
}

func (a *Rest) Start(ctx context.Context) {
	a.server = echo.New()
	a.handlers.Inject(a.server)

	_ = a.server.Start(fmt.Sprintf(":%s", a.port))
}

func (a *Rest) Shutdown(ctx context.Context) {
	err := a.server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Info("Rest server shutdown")
}

type Handlers struct {
	Todos interfaces.TodosRestHandler
}

func NewHandlers(
	todos interfaces.TodosRestHandler,
) *Handlers {
	return &Handlers{
		Todos: todos,
	}
}

func (h *Handlers) Inject(e *echo.Echo) {
	val := reflect.ValueOf(h).Elem()

	group := e.Group("")

	for i := 0; i < val.NumField(); i++ {
		if !val.Type().Field(i).IsExported() {
			continue
		}

		field := val.Field(i).Interface()

		if handler, ok := field.(interfaces.Route); ok {
			handler.RegisterRoute(group)
		}
	}
}
