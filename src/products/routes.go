package products

import "github.com/labstack/echo/v4"

type Routes struct {
	Delivery IDelivery
}

func NewRoutes(delivery IDelivery) *Routes {
	return &Routes{Delivery: delivery}
}

func (routes *Routes) RegisterRoutes(r *echo.Echo) {
	group := r.Group("api/v1")

	group.GET("/products", routes.Delivery.GetListData)
	group.GET("/products/:id", routes.Delivery.GetByID)
	group.POST("/products", routes.Delivery.InsertData)
	group.PUT("/products/:id", routes.Delivery.UpdateDataByID)
	group.DELETE("/products/:id", routes.Delivery.DeleteDataByID)

	// Gemini
	group.POST("/gemini", routes.Delivery.GetContent)
}
