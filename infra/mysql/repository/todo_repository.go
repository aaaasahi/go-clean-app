package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	todoDomain "go-clean-app/domain/todo"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

var _ todoDomain.TodoRepository = new(TodoRepository)

func (tr *TodoRepository) Save(ctx context.Context, todo *todoDomain.Todo) (int64, error) {
	query := `
		INSERT INTO todos (title, description, created_at, updated_at)
		VALUES (?, ?, ?, ?)
	`

	result, err := tr.db.ExecContext(ctx, query, todo.Title(), todo.Description(), todo.CreatedAt(), todo.UpdatedAt())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TodoRepository) FindById(ctx context.Context, id int64) (*todoDomain.Todo, error) {
	query := `
		SELECT id, title, description, created_at, updated_at
		FROM todos
		WHERE id = ?
	`

	row := tr.db.QueryRowContext(ctx, query, id)

	var todoData struct {
		ID          int64
		Title       string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	err := row.Scan(&todoData.ID, &todoData.Title, &todoData.Description, &todoData.CreatedAt, &todoData.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, todoDomain.ErrTodoNotFound
		}
		return nil, err
	}

	return todoDomain.ReConstruct(todoData.ID, todoData.Title, todoData.Description, todoData.CreatedAt, todoData.UpdatedAt)
}