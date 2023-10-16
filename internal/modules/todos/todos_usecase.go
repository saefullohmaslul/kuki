package todos

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type useCase struct {
	todosRepository interfaces.TodosRepository
}

func NewUseCase(todosRepository interfaces.TodosRepository) interfaces.TodosUseCase {
	return &useCase{
		todosRepository: todosRepository,
	}
}

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
}
