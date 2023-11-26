package models

// Todos is struct to represent todos table
type Todos struct {
	Base
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	Completed   bool   `json:"completed" validate:"omitempty"`
}

func (m *Todos) IsExist() bool {
	return m.ID.String() != ""
}
