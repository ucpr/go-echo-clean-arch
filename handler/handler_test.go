package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	test_json_data1 = `{"text":"HelloWorld!"}`
	test_form_data = `text=HelloWorld!`
	test_json_data2 = `{"txt":"hoge"}`
)

func TestIndexHandler(t *testing.T) {
	e := echo.New()

	// post json data

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test_json_data1))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, IndexHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, test_json_data1, strings.TrimRight(rec.Body.String(), "\n"))
	}

	// post form data

	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test_form_data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, IndexHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, test_json_data1, strings.TrimRight(rec.Body.String(), "\n"))
	}

	// post 

	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test_form_data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, IndexHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, test_json_data1, strings.TrimRight(rec.Body.String(), "\n"))
	}

}
