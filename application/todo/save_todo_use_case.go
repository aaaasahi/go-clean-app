package todo

import (
	"context"
	todoDomain "go-clean-app/domain/todo"
	"time"
)

type SaveTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewSaveTodoUseCase(todoRepo todoDomain.TodoRepository) *SaveTodoUseCase {
	return &SaveTodoUseCase{
		todoRepo: todoRepo,
	}
}

type SaveTodoUseCaseDto struct {
	Title       string
	Description string
}

func (tc *SaveTodoUseCase) Run(ctx context.Context, input SaveTodoUseCaseDto) (int64, error) {
	now := time.Now()
	todo, err := todoDomain.NewTodo(input.Title, input.Description, now, now)
	if err != nil {
		return 0, err
	}
	id, err := tc.todoRepo.Save(ctx, todo)
	if err != nil {
		return 0, err
	}
	return id, nil
}
