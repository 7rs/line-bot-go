package main

import (
	"log"
	"net/http"
	"os"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/7rs/line-bot-go/line"
)

var client = line.NewMessagingAPIClient(os.Getenv("CHANNEL_ACCESS_TOKEN"), os.Getenv("CHANNEL_SECRET"))

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

func linebot(c echo.Context, req *http.Request, body []byte) error {
	return line.CreateTestResponse(c, http.StatusOK, "unko")
}

func startEcho() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println("Request Body:", string(reqBody))
	}))

	e.GET("/", index())
	e.POST("/linebot", client.GetHandler(linebot))

	err := e.Start(":" + getPort())
	if err != nil {
		return err
	}
	return nil
}

func main() {
	setColog()

	if err := startEcho(); err != nil {
		log.Fatal(err)
	}
}
