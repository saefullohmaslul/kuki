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

// GetTodo is function to get todo detail
func (r *repository) GetTodo(ctx context.Context, params *dtos.GetTodoRequest) (data dtos.GetTodoResponse, err error) {
	err = r.postgres.DB.
		WithContext(ctx).
		Table("todos").
		Where("todo_id = ?", params.TodoID).
		Find(&data.Todos).
		Error

	return
}

// CreateTodo is function to create todo
func (r *repository) CreateTodo(ctx context.Context, params *models.Todos) (data dtos.CreateTodoResponse, err error) {
	err = r.postgres.DB.
		WithContext(ctx).
		Table("todos").
		Create(&params).
		Error

	if err != nil {
		return
	}

	data.Todos = *params
	return
}

// UpdateTodo is function to update todo
func (r *repository) UpdateTodo(ctx context.Context, params *dtos.UpdateTodoRequest) (data dtos.UpdateTodoResponse, err error) {
	err = r.postgres.DB.
		WithContext(ctx).
		Table("todos").
		Where("todo_id = ?", params.TodoID).
		Updates(&params.Todos).
		Error

	data.Todos = params.Todos
	return
}
