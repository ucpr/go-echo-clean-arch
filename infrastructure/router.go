package infrastructure

import (
	"github.com/labstack/echo"

	"../interfaces/controllers"
)

func init() {
	router := echo.New()
	router.POST("/", func(c echo.Context) (err error) {
		return controllers.IndexHandler(c)
	})


	router.Logger.Fatal(router.Start(":8080"))
}
