package todos

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/constants"
	"github.com/saefullohmaslul/kuki/internal/dtos"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
	"github.com/saefullohmaslul/kuki/internal/pkg/response"
	"github.com/saefullohmaslul/kuki/internal/pkg/validator"
	"net/http"
)

type useCase struct {
	todosRepository interfaces.TodosRepository
	validator       *validator.Validator
}

func NewUseCase(todosRepository interfaces.TodosRepository) interfaces.TodosUseCase {
	return &useCase{
		todosRepository: todosRepository,
	}
}

func (s *useCase) InsertTodo(request *models.Todos) error {
	return s.todosRepository.InsertTodo(request)
}

// DeleteTodoById implements interfaces.TodosUseCase.
func (s *useCase) DeleteTodoById(id string) error {
	return s.todosRepository.DeleteTodoById(id)
}

// FindTodoById implements interfaces.TodosUseCase.
func (s *useCase) FindTodoById(id string) (*models.Todos, error) {
	return s.todosRepository.FindTodoById(id)
}

// UpdateTodoById implements interfaces.TodosUseCase.
func (s *useCase) UpdateTodoById(id string, request *models.Todos) error {
	return s.todosRepository.UpdateTodoById(id, request)
}

func (s *useCase) GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error) {
	if err = s.validator.Validate(params); err != nil {
		err = response.New[dtos.GetTodoResponse]().
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrFailedGetTodo).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	data, err = s.todosRepository.GetTodo(ctx, params)
	if err != nil {
		err = response.New[dtos.GetTodoResponse]().
			Error().
			SetCode(http.StatusInternalServerError).
			SetMessage(constants.ErrFailedGetTodo).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	return
}
