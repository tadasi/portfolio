package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"github.com/tadasi/portfolio/domain/models"
)

type TodoController struct {
	todoRepository models.TodoRepository
}

func NewTodoController(i *do.Injector) (*TodoController, error) {
	return &TodoController{
		todoRepository: do.MustInvoke[models.TodoRepository](i),
	}, nil
}

func (c *TodoController) Mount(group *echo.Group) {
	group.GET("/:id", c.Show)
}

func (c *TodoController) Show(ec echo.Context) error {
	response, err := c.todoRepository.Find(
		ec.Request().Context(),
		ec.Param("id"),
	)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, response)
}
