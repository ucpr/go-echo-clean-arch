package infrastructure

import (
	"net/http"
	"html/template"
	"github.com/labstack/echo"

	"github.com/ucpr/go-echo-sample/interfaces/controllers"
)

func init() {
	router := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	router.Renderer = renderer

	router.POST("/", func(c echo.Context) (err error) {
		return controllers.IndexHandler(c)
	})
	router.GET("/upload", func(c echo.Context) (err error) {
		return c.Render(http.StatusOK, "upload.html", nil)
	})
	router.POST("/upload", func(c echo.Context) (err error) {
		return controllers.UploadHandler(c)
	})


	router.Logger.Fatal(router.Start(":8080"))
}
