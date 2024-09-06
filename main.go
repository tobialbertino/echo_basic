package main

import (
	"echo_basic/app"
	"echo_basic/infra/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		log.Println(err)
	}
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://127.0.0.1:5173", "http://localhost:5173"},
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
	)

	app.InitApp(e, conf)

	e.Logger.Fatal(e.Start(":8000"))
}
