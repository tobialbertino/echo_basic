package main

import (
	"echo_basic/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	app.InitApp(e)

	e.Logger.Fatal(e.Start(":8000"))
}
