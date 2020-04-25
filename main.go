package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8000"
	}
	return port
}

func mainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	}
}

func messagingAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", mainPage())
	e.POST("/linebot", messagingAPI())

	err := e.Start(":" + getPort())
	if err != nil {
		log.Fatal(err)
	}
}
