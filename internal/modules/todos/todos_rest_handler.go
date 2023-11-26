package todos

import (
	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/kuki/internal/constants"
	"github.com/saefullohmaslul/kuki/internal/dtos"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/pkg/response"
	"net/http"
)

type restHandler struct {
	todosUseCase interfaces.TodosUseCase
}

func NewRestHandler(todosUseCase interfaces.TodosUseCase) interfaces.TodosRestHandler {
	return &restHandler{
		todosUseCase: todosUseCase,
	}
}

// RegisterRoute is a function to register route
func (r *restHandler) RegisterRoute(e *echo.Group) {
	todos := e.Group("todos")

	todos.GET("/:todo_id", r.GetTodo)
	todos.POST("/", r.CreateTodo)
	todos.PUT("/:todo_id", r.UpdateTodo)
	todos.DELETE("/:todo_id", r.DeleteTodo)
}

func (r *restHandler) GetTodo(c echo.Context) (err error) {
	var (
		params = dtos.GetTodoRequest{}
		data   = dtos.GetTodoResponse{}
		ctx    = c.Request().Context()
		resp   = response.New[dtos.GetTodoResponse]()
	)

	if err := c.Bind(&params); err != nil {
		err = resp.
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrInvalidPayload).
			SetDetail(err).
			SetContext(ctx)
		return c.JSON(resp.Send(data, err))
	}

	data, err = r.todosUseCase.GetTodo(ctx, &params)
	return c.JSON(resp.Send(data, err))
}

func (r *restHandler) CreateTodo(c echo.Context) (err error) {
	var (
		params = dtos.CreateTodoRequest{}
		data   = dtos.CreateTodoResponse{}
		ctx    = c.Request().Context()
		resp   = response.New[dtos.CreateTodoResponse]()
	)

	if err := c.Bind(&params); err != nil {
		err = resp.
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrInvalidPayload).
			SetDetail(err).
			SetContext(ctx)
		return c.JSON(resp.Send(data, err))
	}

	data, err = r.todosUseCase.CreateTodo(ctx, &params)
	return c.JSON(resp.Send(data, err))
}

func (r *restHandler) UpdateTodo(c echo.Context) error {
	var (
		params = dtos.UpdateTodoRequest{}
		data   = dtos.UpdateTodoResponse{}
		ctx    = c.Request().Context()
		resp   = response.New[dtos.UpdateTodoResponse]()
	)

	if err := c.Bind(&params); err != nil {
		err = resp.
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrInvalidPayload).
			SetDetail(err).
			SetContext(ctx)
		return c.JSON(resp.Send(data, err))
	}

	data, err := r.todosUseCase.UpdateTodo(ctx, &params)
	return c.JSON(resp.Send(data, err))
}

func (r *restHandler) DeleteTodo(c echo.Context) error {
	var (
		params = dtos.DeleteTodoRequest{}
		ctx    = c.Request().Context()
		resp   = response.New[interface{}]()
	)

	if err := c.Bind(&params); err != nil {
		err = resp.
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrInvalidPayload).
			SetDetail(err).
			SetContext(ctx)
		return c.JSON(resp.Send(nil, err))
	}

	err := r.todosUseCase.DeleteTodo(ctx, &params)
	return c.JSON(resp.Send(nil, err))
}
