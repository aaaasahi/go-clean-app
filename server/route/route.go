package route

import (
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	v1 := e.Group("/v1")

	todoRoute(v1)
}

func todoRoute(g *echo.Group) {
	g.GET("/todo", func(c echo.Context) error {
		return c.JSON(200, "TODO")
	})
}