//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package todo

import "context"

type TodoRepository interface {
	Save(ctx context.Context, todo *Todo) (int64, error)
	FindById(ctx context.Context, id int64) (*Todo, error)
}
