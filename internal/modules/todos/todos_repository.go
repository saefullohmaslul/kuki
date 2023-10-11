package todos

import (
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
	"github.com/saefullohmaslul/kuki/internal/pkg/database"
)

type Repository struct {
	database.Database
}

func NewTodosRepository(database database.Database) interfaces.TodosRepository {
	return &Repository{
		Database: database,
	}
}

func (r *Repository) InsertTodo(request *models.Todo) error {
	tx := r.DB.Begin()

	err := tx.Model(&models.Todo{}).Create(request).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}


// DeleteTodoById implements interfaces.TodosRepository.
func (*Repository) DeleteTodoById(id string) error {
	panic("unimplemented")
}

// FindTodoById implements interfaces.TodosRepository.
func (*Repository) FindTodoById(id string) (*models.Todo, error) {
	panic("unimplemented")
}

// UpdateTodoById implements interfaces.TodosRepository.
func (*Repository) UpdateTodoById(id string, request *models.Todo) error {
	panic("unimplemented")
}