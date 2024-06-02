package todo

import "time"

type getTodoResponse struct {
	Todo todoResponseModel `json:"todo"`
}

type todoResponseModel struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type createTodoResponse struct {
	ID int64 `json:"id"`
}
