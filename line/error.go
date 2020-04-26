package line

import (
	"github.com/labstack/echo"
)

type MessagingAPIErrorJSON struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type MessagingAPIError struct{}

func (e *MessagingAPIError) Error() string {
	return "Error occurred in Messaging API."
}

func DoMessagingAPIError(c echo.Context, msg string, code int) (*MessagingAPIErrorJSON, error) {
	return &MessagingAPIErrorJSON{
		Message: msg,
		Code:    code,
	}, &MessagingAPIError{}
}
