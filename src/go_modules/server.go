package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	echos := echo.New()
	echos.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	echos.Logger.Fatal(echos.Start(":1323"))
}
