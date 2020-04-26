package line

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type Client struct {
	Token  string
	Secret string
}

type TestResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func CreateTestResponse(c echo.Context, code int, msg string) error {
	return c.JSON(code, &TestResponse{
		Message: msg,
		Code:    code,
	})
}

func NewMessagingAPIClient(token string, secret string) *Client {
	return &Client{
		Token:  token,
		Secret: secret,
	}
}

func (p *Client) GetHandler(f func(c echo.Context, req *http.Request, body []byte) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errJSON, _ := DoMessagingAPIError(c, "Could not read request body.", http.StatusBadRequest)
			return c.JSON(errJSON.Code, errJSON)
		}

		errJSON, err := p.verifySignature(c, req, body)
		if err != nil {
			return c.JSON(errJSON.Code, errJSON)
		}

		return f(c, req, body)
	}
}
