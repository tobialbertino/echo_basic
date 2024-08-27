package fruits

import "github.com/labstack/echo/v4"

type Routes struct {
	Delivery IDelivery
}

func NewRoutes(delivery IDelivery) *Routes {
	return &Routes{Delivery: delivery}
}

func (routes *Routes) RegisterRoutes(e *echo.Echo) {
	group := e.Group("api/v1")

	group.GET("/fruits", routes.Delivery.GetAllData)
	group.GET("/fruits/:id", routes.Delivery.GetDataByID)
	group.POST("/fruits", routes.Delivery.InsertData)
	group.DELETE("/fruits/:id", routes.Delivery.DeleteDataByID)
}
