package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setColog() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
}

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

func verificationSignature(c echo.Context, req *http.Request, body []byte) error {
	// Get line signature
	receivedSignature, err := base64.StdEncoding.DecodeString(req.Header.Get("X-Line-Signature"))
	if err != nil {
		return err
	}

	// Get hash digest
	hash := hmac.New(sha256.New, []byte(os.Getenv("CHANNEL_SECRET")))
	_, err = hash.Write(body)
	if err != nil {
		return err
	}
	signature := hash.Sum(nil)

	// Verify signature
	log.Println("info: Received:", receivedSignature)
	log.Println("info: Correct:", signature)
	if !hmac.Equal(receivedSignature, signature) {
		return c.String(http.StatusForbidden, "X-Line-Signature is invalid.")
	}

	return nil
}

func messagingAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get request body
		req := c.Request()

		// Read request body
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		// Verify signature
		if err = verificationSignature(c, req, body); err != nil {
			return err
		}

		// Return response
		return c.String(http.StatusOK, "OK")
	}
}

func main() {
	setColog()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println("Request Body:", string(reqBody))
	}))

	e.GET("/", mainPage())
	e.POST("/linebot", messagingAPI())

	err := e.Start(":" + getPort())
	if err != nil {
		log.Fatal(err)
	}
}
