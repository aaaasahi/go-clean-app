package todo

import (
	"errors"
	"time"
)

type Todo struct {
	id          int64
	title       string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

const (
	maxTitleLength       = 50
	maxDescriptionLength = 300
)

func NewTodo(id int64, title, description string, createdAt, updatedAt time.Time) (*Todo, error) {
	return newTodo(id, title, description, createdAt, updatedAt)
}

func newTodo(id int64, title, description string, createdAt, updatedAt time.Time) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if len(title) > maxTitleLength {
		return nil, errors.New("title cannot exceed 50 characters")
	}

	if len(description) > maxDescriptionLength {
		return nil, errors.New("description cannot exceed 300 characters")
	}

	return &Todo{
		id,
		title,
		description,
		createdAt,
		updatedAt,
	}, nil
}

func (t *Todo) ID() int64 {
	return t.id
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Description() string {
	return t.description
}

func (t *Todo) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Todo) UpdatedAt() time.Time {
	return t.updatedAt
}