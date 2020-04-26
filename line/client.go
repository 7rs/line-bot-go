package line

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type APIClient struct {
	Token  string
	Secret string
}

type RequestBodyJSON struct {
	Destination string                   `json:"destination"`
	Events      []map[string]interface{} `json:"events"`
}

func GetJSONResponse(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{})
}

func NewMessagingAPIClient(token string, secret string) *APIClient {
	return &APIClient{
		Token:  token,
		Secret: secret,
	}
}

func (p *APIClient) GetHandler(f func(c echo.Context, req *http.Request, body []byte) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			errJSON, _ := doMessagingAPIError(c, "Could not read request body.", http.StatusBadRequest)
			return c.JSON(errJSON.Code, errJSON)
		}

		errJSON, err := p.verifySignature(c, req, body)
		if err != nil {
			return c.JSON(errJSON.Code, errJSON)
		}

		return f(c, req, body)
	}
}
