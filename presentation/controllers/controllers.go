package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type Controller struct {
	todoController *TodoController
}

func NewController(i *do.Injector) (*Controller, error) {
	return &Controller{
		todoController: do.MustInvoke[*TodoController](i),
	}, nil
}

func (c *Controller) Mount(group *echo.Group) {
	c.todoController.Mount(group.Group("/todos"))
}
