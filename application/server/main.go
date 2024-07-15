package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/samber/do"
	"github.com/tadasi/portfolio/application"
	"github.com/tadasi/portfolio/presentation/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ec := echo.New()
	ec.Validator = application.NewValidator()

	ec.Use(middleware.Logger())
	ec.Use(middleware.Recover())

	injector := application.RegisterInjector()
	controller := do.MustInvoke[*controllers.Controller](injector)
	controller.Mount(ec.Group(""))

	ec.Logger.Fatal(ec.Start(":8080"))
}
