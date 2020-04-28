package main

import (
	"log"
	"net/http"
	"os"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	sdk "github.com/line/line-bot-sdk-go/linebot"
)

var err error
var bot *sdk.Client

func setColog() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Lshortfile,
	})
	colog.Register()
}

func getPort() string {
	port := os.Getenv("PORT")
	if port := os.Getenv("PORT"); port == "" {
		return "8000"
	}
	return port
}

func index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	}
}

func linebot() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		events, err := bot.ParseRequest(req)
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
					if _, err := bot.ReplyMessage(event.ReplyToken, sdk.NewTextMessage(msg.Text)).Do(); err != nil {
						log.Printf("error: %v", err)
					}
				}
			}
		}

		return c.String(http.StatusOK, "OK")
	}
}

func startEcho() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Printf("debug: Request Body: %v", string(reqBody))
	}))

	e.GET("/", index())
	e.POST("/linebot", linebot())

	err := e.Start(":" + getPort())
	if err != nil {
		return err
	}
	return nil
}

func main() {
	bot, err = sdk.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	setColog()

	if err := startEcho(); err != nil {
		log.Fatal(err)
	}
}
