package todo

import (
	"context"
	todoDomain "go-clean-app/domain/todo"
	"time"
)

type SaveTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewSaveTodoUseCase(todoRepo todoDomain.TodoRepository) *FindTodoUseCase {
	return &FindTodoUseCase{
		todoRepo: todoRepo,
	}
}

type SaveTodoUseCaseDto struct {
	Title       string
	Description string
}

func (tc *SaveTodoUseCase) Run(ctx context.Context, input SaveTodoUseCaseDto) error {
	now := time.Now()
	todo, err := todoDomain.NewTodo(input.Title, input.Description, now, now)
	if err != nil {
		return err
	}
	return tc.todoRepo.Save(ctx, todo)
}
