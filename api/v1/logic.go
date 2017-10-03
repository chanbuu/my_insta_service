package v1

import "github.com/labstack/echo"
import "net/http"

func TFunc(c echo.Context) error {
	return c.String(http.StatusOK, "OK. Access")

}
