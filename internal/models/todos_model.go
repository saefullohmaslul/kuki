package models

// Todos is struct to represent todos table
type Todos struct {
	TodoID      string `json:"todo_id" gorm:"primaryKey"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed,omitempty"`
}
