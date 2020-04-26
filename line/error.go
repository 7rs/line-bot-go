package line

import (
	"github.com/labstack/echo"
)

type messagingAPIErrorJSON struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type messagingAPIError struct{}

func (e *messagingAPIError) Error() string {
	return "Error occurred in Messaging API."
}

func doMessagingAPIError(c echo.Context, msg string, code int) (*messagingAPIErrorJSON, error) {
	return &messagingAPIErrorJSON{
		Message: msg,
		Code:    code,
	}, &messagingAPIError{}
}
