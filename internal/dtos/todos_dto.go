package dtos

import "github.com/saefullohmaslul/kuki/internal/models"

type (
	GetTodoRequest struct {
		TodoID string `json:"todo_id" validate:"required"`
	}

	GetTodoResponse struct {
		TodoID      string `json:"todo_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
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
