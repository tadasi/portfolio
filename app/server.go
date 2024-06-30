package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ec := echo.New()

	ec.Use(middleware.Logger())
	ec.Use(middleware.Recover())

	ec.GET("/", hello)

	ec.Logger.Fatal(ec.Start(":8080"))
}

func hello(ec echo.Context) error {
	return ec.String(http.StatusOK, "Hello, World!")
}
