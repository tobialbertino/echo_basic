package words

import "github.com/labstack/echo/v4"

type Routes struct {
	Delivery IDelivery
}

func NewRoutes(delivery IDelivery) *Routes {
	return &Routes{Delivery: delivery}
}

func (routes *Routes) RegisterRoutes(e *echo.Echo) {
	group := e.Group("api/v1")

	group.GET("/words", routes.Delivery.GetAllData)
	group.POST("/words", routes.Delivery.InsertData)
}
