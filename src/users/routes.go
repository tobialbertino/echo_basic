package users

import "github.com/labstack/echo/v4"

type Routes struct {
	Delivery IDelivery
}

func NewRoutes(delivery IDelivery) *Routes {
	return &Routes{Delivery: delivery}
}

func (routes *Routes) RegisterRoutes(r *echo.Echo) {
	group := r.Group("api/v1")

	group.GET("/users", routes.Delivery.GetListData)
	group.GET("/users/:id", routes.Delivery.GetByID)
	group.POST("/users", routes.Delivery.InsertData)
	group.PUT("/users/:id", routes.Delivery.UpdateDataByID)
	group.DELETE("/users/:id", routes.Delivery.DeleteDataByID)
}
