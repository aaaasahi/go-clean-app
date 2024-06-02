package di

import (
	"database/sql"
	todoApp "go-clean-app/application/todo"
	todoDomain "go-clean-app/domain/todo"
	repository "go-clean-app/infra/mysql/repository"
	handler "go-clean-app/presentation/todo"
	"go-clean-app/server"
	"go-clean-app/server/route"

	"go.uber.org/dig"
)

func BuildContainer(db *sql.DB) *dig.Container {
	container := dig.New()

	container.Provide(func() todoDomain.TodoRepository {
		return repository.NewTodoRepository(db)
	})
	container.Provide(todoApp.NewFindTodoUseCase)
	container.Provide(todoApp.NewSaveTodoUseCase)
	container.Provide(handler.NewTodoHandler)
	container.Provide(route.NewRouter)
	container.Provide(server.NewServer)

	return container
}