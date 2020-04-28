package bot

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (b *Bot) SetEcho() {
	b.Echo.Use(middleware.Logger())
	b.Echo.Use(middleware.Recover())
	b.Echo.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Printf("debug: Request Body: %v", string(reqBody))
	}))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port := os.Getenv("PORT"); port == "" {
		return "8000"
	}
	return port
}

func (b *Bot) SetEndpoint(method string, path string, e echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	switch method {
	case http.MethodGet:
		b.Echo.GET(path, e)
	case http.MethodHead:
		b.Echo.HEAD(path, e)
	case http.MethodPost:
		b.Echo.POST(path, e)
	case http.MethodPut:
		b.Echo.PUT(path, e)
	case http.MethodPatch:
		b.Echo.PATCH(path, e)
	case http.MethodDelete:
		b.Echo.DELETE(path, e)
	case http.MethodConnect:
		b.Echo.CONNECT(path, e)
	case http.MethodOptions:
		b.Echo.OPTIONS(path, e)
	case http.MethodTrace:
		b.Echo.TRACE(path, e)
	}
}

func (b *Bot) StartEcho() error {
	if err := b.Echo.Start(":" + getPort()); err != nil {
		return err
	}
	return nil
}
