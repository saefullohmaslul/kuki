package models

type Todos struct {
	TodoID      string `json:"todo_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
