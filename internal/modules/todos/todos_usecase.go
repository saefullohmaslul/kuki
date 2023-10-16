package todos

import (
	"context"
<<<<<<< HEAD

	"github.com/opentracing/opentracing-go"
=======
	"github.com/saefullohmaslul/kuki/internal/constants"
	"github.com/saefullohmaslul/kuki/internal/dtos"
>>>>>>> main
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/pkg/response"
	"github.com/saefullohmaslul/kuki/internal/pkg/validator"
	"net/http"
)

type useCase struct {
	todosRepository interfaces.TodosRepository
	validator       *validator.Validator
}

// NewUseCase is function to create new todos use case
func NewUseCase(todosRepository interfaces.TodosRepository) interfaces.TodosUseCase {
	return &useCase{
		todosRepository: todosRepository,
	}
}

<<<<<<< HEAD
func (s *useCase) InsertTodo(ctx context.Context, request *models.Todos) error {
	trace, _ := opentracing.StartSpanFromContext(ctx, "useCase.InsertTodo")
	defer trace.Finish()
	return s.todosRepository.InsertTodo(request)
}

// DeleteTodoById implements interfaces.TodosUseCase.
func (s *useCase) DeleteTodoById(ctx context.Context, id string) error {
	todo, err := s.FindTodoById(ctx, id)
	if err != nil {
		return err
	}
	trace, _ := opentracing.StartSpanFromContext(ctx, "useCase.DeleteTodoById")
	defer trace.Finish()
	return s.todosRepository.DeleteTodoById(todo.TodoID)
}

// FindTodoById implements interfaces.TodosUseCase.
func (s *useCase) FindTodoById(ctx context.Context, id string) (*models.Todos, error) {
	trace, _ := opentracing.StartSpanFromContext(ctx, "useCase.FindTodoById")
	defer trace.Finish()
	return s.todosRepository.FindTodoById(id)
}

// UpdateTodoById implements interfaces.TodosUseCase.
func (s *useCase) UpdateTodoById(ctx context.Context, id string, request *models.Todos) error {
	todo, err := s.FindTodoById(ctx, id)
	if err != nil {
		return err
	}
	trace, _ := opentracing.StartSpanFromContext(ctx, "useCase.UpdateTodoById")
	defer trace.Finish()
	return s.todosRepository.UpdateTodoById(todo.TodoID, request)
=======
// GetTodo is function to get todo detail
func (s *useCase) GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error) {
	if err = s.validator.Validate(params); err != nil {
		err = response.New[dtos.GetTodoResponse]().
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrFailedValidateRequest).
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

// CreateTodo is function to create todo
func (s *useCase) CreateTodo(ctx context.Context, params *dtos.CreateTodoRequest) (data dtos.CreateTodoResponse, err error) {
	if err = s.validator.Validate(params); err != nil {
		err = response.New[dtos.CreateTodoResponse]().
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrFailedValidateRequest).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	data, err = s.todosRepository.CreateTodo(ctx, &params.Todos)
	if err != nil {
		err = response.New[dtos.CreateTodoResponse]().
			Error().
			SetCode(http.StatusInternalServerError).
			SetMessage(constants.ErrFailedCreateTodo).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	return
}

// UpdateTodo is function to update todo
func (s *useCase) UpdateTodo(ctx context.Context, params *dtos.UpdateTodoRequest) (data dtos.UpdateTodoResponse, err error) {
	if err = s.validator.Validate(params); err != nil {
		err = response.New[dtos.UpdateTodoResponse]().
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrFailedValidateRequest).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	data, err = s.todosRepository.UpdateTodo(ctx, params)
	if err != nil {
		err = response.New[dtos.UpdateTodoResponse]().
			Error().
			SetCode(http.StatusInternalServerError).
			SetMessage(constants.ErrFailedUpdateTodo).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	return
}

// DeleteTodo is function to delete todo
func (s *useCase) DeleteTodo(ctx context.Context, params *dtos.DeleteTodoRequest) (err error) {
	if err = s.validator.Validate(params); err != nil {
		err = response.New[dtos.GetTodoResponse]().
			Error().
			SetCode(http.StatusBadRequest).
			SetMessage(constants.ErrFailedValidateRequest).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	err = s.todosRepository.DeleteTodo(ctx, params)
	if err != nil {
		err = response.New[dtos.GetTodoResponse]().
			Error().
			SetCode(http.StatusInternalServerError).
			SetMessage(constants.ErrFailedDeleteTodo).
			SetDetail(err).
			SetContext(ctx)
		return
	}

	return
>>>>>>> main
}
