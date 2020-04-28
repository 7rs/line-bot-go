package bot

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	sdk "github.com/line/line-bot-sdk-go/linebot"
)

func (b *Bot) SetEndPoints() {
	b.SetEndpoint(http.MethodGet, "/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
	b.SetEndpoint(http.MethodPost, "/linebot", func(c echo.Context) error {
		req := c.Request()
		events, err := b.API.ParseRequest(req)
		if err != nil {
			if err == sdk.ErrInvalidSignature {
				return c.String(http.StatusBadRequest, "Bad Request")
			} else {
				return c.String(http.StatusInternalServerError, "Internal Server Error")
			}
		}

		for _, event := range events {
			if event.Type == sdk.EventTypeMessage {
				switch msg := event.Message.(type) {
				case *sdk.TextMessage:
					if _, err := b.API.ReplyMessage(event.ReplyToken, sdk.NewTextMessage(msg.Text)).Do(); err != nil {
						log.Printf("error: %v", err)
					}
				}
			}
		}

		return c.String(http.StatusOK, "OK")
	})
}
