package controllers

type Context interface {
	JSON(code int, body interface{}) error
	Bind(body interface{}) error
}
