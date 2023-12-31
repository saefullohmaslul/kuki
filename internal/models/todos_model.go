package models

// Todos is struct to represent todos table
type Todos struct {
	TodoID      string `json:"todo_id;primaryKey"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	Completed   bool   `json:"completed" validate:"omitempty"`
}
