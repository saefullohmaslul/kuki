package todos

import (
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type UseCase struct {
	todosRepository interfaces.TodosRepository
}

func NewUseCase(todosRepository interfaces.TodosRepository) interfaces.TodosUseCase {
	return &UseCase{
		todosRepository: todosRepository,
	}
}

func (s *UseCase) InsertTodo(request *models.Todos) error {
	return s.todosRepository.InsertTodo(request)
}

// DeleteTodoById implements interfaces.TodosUseCase.
func (s *UseCase) DeleteTodoById(id string) error {
	return s.todosRepository.DeleteTodoById(id)
}

// FindTodoById implements interfaces.TodosUseCase.
func (s *UseCase) FindTodoById(id string) (*models.Todos, error) {
	return s.todosRepository.FindTodoById(id)
}

// UpdateTodoById implements interfaces.TodosUseCase.
func (s *UseCase) UpdateTodoById(id string, request *models.Todos) error {
	return s.todosRepository.UpdateTodoById(id, request)
}
