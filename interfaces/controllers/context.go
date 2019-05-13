package controllers

import (
	"mime/multipart"
)

type Context interface {
	JSON(code int, body interface{}) error
	Bind(body interface{}) error
	MultipartForm() (*multipart.Form, error)
	String(code int, body string) error
}
