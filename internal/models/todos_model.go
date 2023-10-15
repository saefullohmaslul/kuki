package models

type Todos struct {
	TodoID      string `json:"todo_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	Completed   bool   `json:"completed" validate:"omitempty"`
}
