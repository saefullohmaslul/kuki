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

func (r *Repository) InsertTodo(request *models.Todo) (*models.Todo, error) {
	tx := r.DB.Begin()

	if err := tx.Model(&models.Todo{}).Create(request).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return request, nil
}
