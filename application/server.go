package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tadasi/portfolio/infrastructure/mysql"
	"github.com/tadasi/portfolio/infrastructure/mysql/tables"

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

// TODO: 後で Domain に移動する
type Todo struct {
	ID          string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

func hoge(ec echo.Context) error {
	ctx := context.Background()
	db, err := mysql.Open()
	if err != nil {
		return err
	}
	todo, err := tables.Todos(
		tables.TodoWhere.ID.EQ("12345"), // TODO: 後で修正する
	).One(ctx, db)
	if err != nil {
		return err
	}

	fmt.Println(`debug/todo`, todo) // TODO: 後で消す
	return ec.String(http.StatusOK, "hoge")
}

func fuga(ec echo.Context) error {
	return ec.String(http.StatusOK, "fuga")
}
