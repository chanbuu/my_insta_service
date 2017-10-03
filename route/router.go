package router

import "github.com/labstack/echo"
import "github.com/my_service/api/v1"
import "net/http"

func NewRouter() *echo.Echo {
	e := echo.New()
	e1 := e.Group("/api/v1")
	e1.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World")
	})
	e1.GET("/access", v1.TFunc)
	return e
}
