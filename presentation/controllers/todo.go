package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"github.com/tadasi/portfolio/application/usecases"
)

type TodoController struct {
	todoInteractor usecases.TodoInteractor
}

func NewTodoController(i *do.Injector) (*TodoController, error) {
	return &TodoController{
		todoInteractor: do.MustInvoke[usecases.TodoInteractor](i),
	}, nil
}

func (c *TodoController) Mount(group *echo.Group) {
	group.GET("/:id", c.Show)
	group.POST("/", c.Create)
}

func (c *TodoController) Show(ec echo.Context) error {
	response, err := c.todoInteractor.FindTodo(
		ec.Request().Context(),
		&usecases.FindTodoInput{
			TodoID: ec.Param("id"),
		},
	)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, response)
}

func (c *TodoController) Create(ec echo.Context) error {
	_, err := c.todoInteractor.CreateTodo(
		ec.Request().Context(),
		&usecases.CreateTodoInput{
			Content: ec.Param("Content"), // TODO: Bind, Validation をかける
		},
	)
	if err != nil {
		return err
	}
	return nil
}
