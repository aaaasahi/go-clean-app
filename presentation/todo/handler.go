package todo

import (
	todoApp "go-clean-app/application/todo"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	findTodoUseCase *todoApp.FindTodoUseCase
	saveTodoUseCase *todoApp.SaveTodoUseCase
}

func NewTodoHandler(findTodoUseCase *todoApp.FindTodoUseCase, saveTodoUseCase *todoApp.SaveTodoUseCase) handler {
	return handler{
		findTodoUseCase: findTodoUseCase,
		saveTodoUseCase: saveTodoUseCase,
	}
}

func (h *handler) GetTodoByID(c echo.Context) error {
	idParam := c.Param("id")
  id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	dto, err := h.findTodoUseCase.Run(c.Request().Context(), id)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	res := getTodoResponse{
		Todo:  todoResponseModel{
			ID:          dto.ID,
			Title:       dto.Title,
			Description: dto.Description,
			CreatedAt:   dto.CreatedAt,
			UpdatedAt:   dto.UpdatedAt,
		},
	}
	return c.JSON(200, res)
}

func (h *handler) SaveTodo(c echo.Context) error {
	var params saveTodoParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(400, err.Error())
	}

	input := todoApp.SaveTodoUseCaseDto{
		Title:       params.Title,
		Description: params.Description,
	}
	id, err := h.saveTodoUseCase.Run(c.Request().Context(), input)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	res := createTodoResponse{
		ID: id,
	}

	return c.JSON(200, res)
}
