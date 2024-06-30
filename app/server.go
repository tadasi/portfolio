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
	ec.GET("/hoge", hoge)
	ec.POST("/fuga", fuga)

	ec.Logger.Fatal(ec.Start(":8080"))
}

func hello(ec echo.Context) error {
	return ec.String(http.StatusOK, "Hello, World!")
}

func hoge(ec echo.Context) error {
	return ec.String(http.StatusOK, "hoge")
}

func fuga(ec echo.Context) error {
	return ec.String(http.StatusOK, "fuga")
}
