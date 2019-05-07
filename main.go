package main

import (
	"net/http"
	"github.com/labstack/echo"
)

type Text struct {
  Body string `json:"text" form:"text" query:"text"`
}

func handler(c echo.Context) (err error) {
  u := new(Text)
  if err = c.Bind(u); err != nil {
    return
  }
  return c.JSON(http.StatusOK, u)
}


func main() {
	e := echo.New()

	e.POST("/", handler)

	e.Logger.Fatal(e.Start(":8080"))
}
