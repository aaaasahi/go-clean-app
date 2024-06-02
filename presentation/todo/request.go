package todo

type saveTodoParams struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}
