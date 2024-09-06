package app

import (
	"echo_basic/infra/config"
	"echo_basic/infra/gemini"
	"echo_basic/src/fruits"
	"echo_basic/src/products"
	"echo_basic/src/users"
	"echo_basic/src/words"
	"github.com/labstack/echo/v4"
)

func InitApp(e *echo.Echo, conf *config.AppConfig) {
	//external Gemini
	geminiService := gemini.NewService(conf)

	// domain users
	usersService := users.NewService()
	usersDelivery := users.NewDelivery(usersService)
	usersRoutes := users.NewRoutes(usersDelivery)
	usersRoutes.RegisterRoutes(e)

	// domain words
	wordsService := words.NewService()
	wordsDelivery := words.NewDelivery(wordsService)
	wordsRoutes := words.NewRoutes(wordsDelivery)
	wordsRoutes.RegisterRoutes(e)

	// domain fruits
	fruitService := fruits.NewService()
	fruitDelivery := fruits.NewDelivery(fruitService)
	fruitsRoutes := fruits.NewRoutes(fruitDelivery)
	fruitsRoutes.RegisterRoutes(e)

	// domain products
	productService := products.NewService(geminiService)
	productDelivery := products.NewDelivery(productService)
	productsRoutes := products.NewRoutes(productDelivery)
	productsRoutes.RegisterRoutes(e)
}
