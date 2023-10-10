package todos

import (
	"context"
	"errors"
	"github.com/saefullohmaslul/kuki/internal/constants"
	"github.com/saefullohmaslul/kuki/internal/grpc"
	"github.com/saefullohmaslul/kuki/internal/pkg/response"
	"net/http"
)

type useCase struct {
}

func NewUsecase() UseCase {
	return &useCase{}
}

func (u *useCase) GetTodo(ctx context.Context, params *grpc.GetTodoRequest) (data *grpc.Todo, err error) {
	// simulate validation
	err = errors.New("some error")
	if err != nil {
		err = response.New[*grpc.Todo]().
			Error().
			SetCode(http.StatusInternalServerError).
			SetMessage(constants.ErrFailedGetTodo).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	data = &grpc.Todo{
		TodoId:      "abc",
		Title:       "Makan siang",
		Description: "Makan siang dengan istri",
		Completed:   false,
	}

	return
}
