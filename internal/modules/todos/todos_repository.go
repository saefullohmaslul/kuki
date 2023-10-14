package todos

import (
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
	"github.com/saefullohmaslul/kuki/internal/pkg/database"
)

type Repository struct {
	postgres database.Postgres
}

func NewRepository(postgres database.Postgres) interfaces.TodosRepository {
	return &Repository{
		postgres: postgres,
	}
}

func (r *Repository) InsertTodo(request *models.Todos) error {
	tx := r.postgres.DB.Begin()

	err := tx.Model(&models.Todos{}).Create(request).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// FindTodoById implements interfaces.TodosRepository.
func (r *Repository) FindTodoById(id string) (*models.Todos, error) {
	var todo models.Todos

	err := r.postgres.DB.Model(&models.Todos{}).Where("todo_id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// UpdateTodoById implements interfaces.TodosRepository.
func (r *Repository) UpdateTodoById(id string, request *models.Todos) error {
	tx := r.postgres.DB.Begin()

	err := tx.Model(&models.Todos{}).Where("todo_id = ?", id).Updates(request).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// DeleteTodoById implements interfaces.TodosRepository.
func (r *Repository) DeleteTodoById(id string) error {
	tx := r.postgres.DB.Begin()

	err := tx.Model(&models.Todos{}).Where("todo_id = ?", id).Delete(&models.Todos{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
