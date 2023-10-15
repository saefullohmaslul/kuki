package dtos

import "github.com/saefullohmaslul/kuki/internal/models"

type (
	GetTodoRequest struct {
		TodoID string `json:"todo_id" validate:"required"`
	}

	GetTodoResponse struct {
		models.Todos
	}
)

type (
	CreateTodoRequest struct {
		models.Todos
	}

	CreateTodoResponse struct {
		models.Todos
	}
)

type (
	UpdateTodoRequest struct {
		models.Todos
	}

	UpdateTodoResponse struct {
		models.Todos
	}
)
