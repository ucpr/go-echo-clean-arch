package controllers

import (
	"net/http"

	"../../domain"
//	"../../usecase"
)

func IndexHandler(c Context) (err error) {
	u := new(domain.Data)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
