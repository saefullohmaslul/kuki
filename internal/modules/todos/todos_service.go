package todos

import (
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
)

type Service struct {
	Repository interfaces.TodosRepository
}

func NewTodosService(todosRepository interfaces.TodosRepository) interfaces.TodosService {
	return &Service{
		Repository: todosRepository,
	}
}

func (s *Service) InsertTodo(request *models.Todo) error {
	return s.Repository.InsertTodo(request)
}

// DeleteTodoById implements interfaces.TodosService.
func (*Service) DeleteTodoById(id string) error {
	panic("unimplemented")
}

// FindTodoById implements interfaces.TodosService.
func (s*Service) FindTodoById(id string) (*models.Todo, error) {
	return s.Repository.FindTodoById(id)
}

// UpdateTodoById implements interfaces.TodosService.
func (s*Service) UpdateTodoById(id string, request *models.Todo) error {
	return s.Repository.UpdateTodoById(id, request)
}
