package line

import "github.com/labstack/echo"

type LineAPIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newLineAPIError(c echo.Context, code int, msg string) error {
	return c.JSON(code, LineAPIError{
		Message: msg,
		Code:    code,
	})
}
