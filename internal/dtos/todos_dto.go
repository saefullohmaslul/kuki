package dtos

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
