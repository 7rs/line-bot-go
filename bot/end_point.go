package bot

import (
	"net/http"

	"github.com/labstack/echo"
)

func index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	}
}

func (p *Client) linebot() echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := p.getEvents(c)
		if err != nil {
			return handleError(c, err)
		}

		p.handleEvents(events)

		return c.String(http.StatusOK, "OK")
	}
}

func (p *Client) SetEndPoints() {
	e.GET("/", index())
	e.POST("/linebot", p.linebot())
}
