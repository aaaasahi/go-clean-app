package todo

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewTodo(t *testing.T) {
	type args struct {
		title       string
		description string
		createdAt   time.Time
		updatedAt   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Todo
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				title:       "title",
				description: "description",
				createdAt:   time.Now(),
				updatedAt:   time.Now(),
			},
			want: &Todo{
				title:       "title",
				description: "description",
				createdAt:   time.Now(),
				updatedAt:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "異常系: titleが空",
			args: args{
				title:       "",
				description: "description",
				createdAt:   time.Now(),
				updatedAt:   time.Now(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: titleが長い",
			args: args{
				title:       "this title is way too long and exceeds the maximum length of 50 characters",
				description: "description",
				createdAt:   time.Now(),
				updatedAt:   time.Now(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: descriptionが長い",
			args: args{
				title:       "title",
				description: "this description is way too long and exceeds the maximum length of 300 characters. " +
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, " +
					"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
				createdAt: time.Now(),
				updatedAt: time.Now(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTodo(tt.args.title, tt.args.description, tt.args.createdAt, tt.args.updatedAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(Todo{}), cmpopts.IgnoreFields(Todo{}, "createdAt", "updatedAt")); diff != "" {
				t.Errorf("NewTodo() = %v, want %v, diff: %s", got, tt.want, diff)
			}
		})
	}
}
