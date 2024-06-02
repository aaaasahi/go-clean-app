package todo

import (
	"context"
	"errors"
	"testing"
	"time"

	todoDomain "go-clean-app/domain/todo"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFindTodoUseCase_Run(t *testing.T) {
  ctrl := gomock.NewController(t)
	mockTodoRepo := todoDomain.NewMockTodoRepository(ctrl)
	uc := NewFindTodoUseCase(mockTodoRepo)

  now := time.Now()

	tests := []struct {
		name     string
		mockFunc func()
		id       int64
		want     *FindTodoUseCaseDto
		wantErr  bool
	}{
		{
			name: "Todoを正常に取得できること",
			mockFunc: func() {
				todo, _ := reconstructTodo(1, "Test Title", "Test Description", now, now)
				mockTodoRepo.EXPECT().FindById(gomock.Any(), int64(1)).Return(todo, nil)
			},
			id: 1,
			want: &FindTodoUseCaseDto{
				ID:          1,
				Title:       "Test Title",
				Description: "Test Description",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
		},
		{
			name: "存在しないTodoを取得するとエラーが返ること",
			mockFunc: func() {
				mockTodoRepo.EXPECT().FindById(gomock.Any(), int64(2)).Return(nil, errors.New("not found"))
			},
			id:      2,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			got, err := uc.Run(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
					t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
					return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func reconstructTodo(id int64, title, description string, createdAt, updatedAt time.Time) (*todoDomain.Todo, error) {
  return todoDomain.ReConstruct(id, title, description, createdAt, updatedAt)
}
