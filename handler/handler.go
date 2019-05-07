package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

type Text struct {
  Body string `json:"text" form:"text" query:"text"`
}

func IndexHandler(c echo.Context) (err error) {
  u := new(Text)
  if err = c.Bind(u); err != nil {
    return
  }
  return c.JSON(http.StatusOK, u)
}

