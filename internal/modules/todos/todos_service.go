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

func (s *Service) InsertTodo(request *models.Todo) (*models.Todo, error) {
	request.TodoId = "1"
	return s.Repository.InsertTodo(request)
}
