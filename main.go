package main

import (
	rest "github.com/jacintojosh/rest"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	registerControllers(e)
	e.Logger.Fatal(e.Start(":9091"))
}

func registerControllers(e *echo.Echo) {
	rest.InitServerHandler(e)
}