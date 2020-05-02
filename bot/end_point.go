package bot

import (
	"net/http"

	"github.com/labstack/echo"
)

func (b *Bot) index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	}
}

func (b *Bot) linebot() echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := b.getEvents(c)
		if err != nil {
			return handleError(c, err)
		}

		b.handleEvents(events)

		return c.String(http.StatusOK, "OK")
	}
}

func (b *Bot) SetEndPoints() {
	b.SetEndpoint(http.MethodGet, "/", b.index())
	b.SetEndpoint(http.MethodPost, "/linebot", b.linebot())
}
