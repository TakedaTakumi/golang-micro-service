package main

import (
	"fmt"
	config "golang-micro-service/configs/command/partner"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, Query Partner!")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
