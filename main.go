package main

import (
	"github.com/labstack/echo"
	"github.com/ucpr/go-echo-sample/handler"
)

func main() {
	e := echo.New()

	e.POST("/", handler.IndexHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
