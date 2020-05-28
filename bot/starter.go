package bot

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetEcho() {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
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

func StartEcho() error {
	if err := e.Start(":" + getPort()); err != nil {
		return err
	}
	return nil
}
