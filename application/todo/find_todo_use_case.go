package todo

import (
	"context"
	todoDomain "go-clean-app/domain/todo"
	"time"
)

type FindTodoUseCase struct {
	todoRepo todoDomain.TodoRepository
}

func NewFindTodoUseCase(todoRepo todoDomain.TodoRepository) *FindTodoUseCase {
	return &FindTodoUseCase{
		todoRepo: todoRepo,
	}
}

type FindTodoUseCaseDto struct {
	ID          int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (tc *FindTodoUseCase) Run(ctx context.Context, id int64) (*FindTodoUseCaseDto, error) {
	todo, err := tc.todoRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindTodoUseCaseDto{
		ID:          todo.ID(),
		Title:       todo.Title(),
		Description: todo.Description(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}, nil
}
