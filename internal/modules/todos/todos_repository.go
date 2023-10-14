package todos

import (
	"context"
	"github.com/saefullohmaslul/kuki/internal/dtos"
	"github.com/saefullohmaslul/kuki/internal/interfaces"
	"github.com/saefullohmaslul/kuki/internal/models"
	"github.com/saefullohmaslul/kuki/internal/pkg/database"
)

type repository struct {
	postgres database.Postgres
}

func NewRepository(postgres database.Postgres) interfaces.TodosRepository {
	return &repository{
		postgres: postgres,
	}
}

func (r *repository) InsertTodo(request *models.Todos) error {
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
func (r *repository) FindTodoById(id string) (*models.Todos, error) {
	var todo models.Todos

	err := r.postgres.DB.Model(&models.Todos{}).Where("todo_id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// UpdateTodoById implements interfaces.TodosRepository.
func (r *repository) UpdateTodoById(id string, request *models.Todos) error {
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
func (r *repository) DeleteTodoById(id string) error {
	tx := r.postgres.DB.Begin()

	err := tx.Model(&models.Todos{}).Where("todo_id = ?", id).Delete(&models.Todos{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *repository) GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error) {
	err = r.postgres.DB.
		WithContext(ctx).
		Table("todos").
		Where("todo_id = ?", params.TodoID).
		Find(&data).
		Error

	return
}
