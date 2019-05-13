package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	testJsonBody = `{"text":"HelloWorld"}`
)

func TestIndexHandler(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testJsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// OK
	if assert.NoError(t, IndexHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"text":"HelloWorld"}`, strings.TrimRight(rec.Body.String(), "\n"))
	}
}
