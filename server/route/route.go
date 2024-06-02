package route

import (
	handler "go-clean-app/presentation/todo"

	"github.com/labstack/echo/v4"
)

type Router struct {
	TodoHandler *handler.Handler
}

func NewRouter(todoHandler *handler.Handler) *Router {
	return &Router{
		TodoHandler: todoHandler,
	}
}

func (r *Router) InitRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	v1 := e.Group("/v1")
	r.todoRoute(v1)
}

func (r *Router) todoRoute(e *echo.Group) {
	group := e.Group("/todos")
	group.GET("/:id", r.TodoHandler.GetTodoByID)
	group.POST("", r.TodoHandler.SaveTodo)
}