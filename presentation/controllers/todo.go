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
	group.POST("/", c.Create)
	group.GET("/:id", c.Show)
	group.PATCH("/:id", c.Update)
	group.DELETE("/:id", c.Delete)
	group.PATCH("/:id/completes", c.Complete)
	group.PATCH("/:id/uncompletes", c.Uncomplete)
}

func (c *TodoController) Create(ec echo.Context) error {
	params := &struct {
		Content string `json:"content" validate:"required"`
	}{}
	if err := ec.Bind(params); err != nil {
		return err
	}
	if err := ec.Validate(params); err != nil {
		return err
	}

	output, err := c.todoInteractor.CreateTodo(
		ec.Request().Context(),
		&usecases.CreateTodoInput{
			Content: params.Content,
		},
	)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusCreated, map[string]string{
		"id": output.TodoID,
	})
}

func (c *TodoController) Show(ec echo.Context) error {
	params := &struct {
		TodoID string `param:"id" validate:"uuid,required"`
	}{}
	if err := ec.Bind(params); err != nil {
		return err
	}
	if err := ec.Validate(params); err != nil {
		return err
	}

	output, err := c.todoInteractor.FindTodo(
		ec.Request().Context(),
		&usecases.FindTodoInput{
			TodoID: params.TodoID,
		},
	)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, output)
}

func (c *TodoController) Update(ec echo.Context) error {
	params := &struct {
		TodoID  string `param:"id" validate:"uuid,required"`
		Content string `json:"content" validate:"required"`
	}{}
	if err := ec.Bind(params); err != nil {
		return err
	}
	if err := ec.Validate(params); err != nil {
		return err
	}

	output, err := c.todoInteractor.UpdateTodo(
		ec.Request().Context(),
		&usecases.UpdateTodoInput{
			TodoID:  params.TodoID,
			Content: params.Content,
		},
	)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, output)
}

func (c *TodoController) Delete(ec echo.Context) error {
	params := &struct {
		TodoID string `param:"id" validate:"uuid,required"`
	}{}
	if err := ec.Bind(params); err != nil {
		return err
	}
	if err := ec.Validate(params); err != nil {
		return err
	}

	if err := c.todoInteractor.DeleteTodo(
		ec.Request().Context(),
		&usecases.DeleteTodoInput{
			TodoID: params.TodoID,
		},
	); err != nil {
		return err
	}
	return ec.JSON(http.StatusNoContent, nil)
}

func (c *TodoController) Complete(ec echo.Context) error {
	params := &struct {
		TodoID string `param:"id" validate:"uuid,required"`
	}{}
	if err := ec.Bind(params); err != nil {
		return err
	}
	if err := ec.Validate(params); err != nil {
		return err
	}

	output, err := c.todoInteractor.CompleteTodo(
		ec.Request().Context(),
		&usecases.CompleteTodoInput{
			TodoID: params.TodoID,
		},
	)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, output)
}

func (c *TodoController) Uncomplete(ec echo.Context) error {
	params := &struct {
		TodoID string `param:"id" validate:"uuid,required"`
	}{}
	if err := ec.Bind(params); err != nil {
		return err
	}
	if err := ec.Validate(params); err != nil {
		return err
	}

	output, err := c.todoInteractor.UncompleteTodo(
		ec.Request().Context(),
		&usecases.UncompleteTodoInput{
			TodoID: params.TodoID,
		},
	)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, output)
}
