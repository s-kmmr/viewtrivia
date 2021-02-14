package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/s-kmmr/viewtrivia/injector"
)

func main() {
	e := echo.New()
	injector.NewInjector(e)

	if err := e.Start(":1323"); err != nil {
		log.Fatal(err)
	}
}
