package main

import (
	"github.com/labstack/echo"
	"./handler"
)

func main() {
	e := echo.New()

	e.POST("/", handler.IndexHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
