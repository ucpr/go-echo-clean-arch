package controllers

import (
	"net/http"

	"github.com/ucpr/go-echo-clean-arch/domain"
//	"../../usecase"
)

func IndexHandler(c Context) (err error) {
	u := new(domain.Data)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
