package fruits

import (
	"context"
	"echo_basic/infra/fruitApi"
	"github.com/labstack/echo/v4"
)

// Domain Users
type IDelivery interface {
	GetAllData(c echo.Context) error
	GetDataByID(c echo.Context) error
	InsertData(c echo.Context) error
	DeleteDataByID(c echo.Context) error
}

type IService interface {
	GetAllData(ctx context.Context) ([]fruitApi.DataFruit, error)
	GetDataByID(ctx context.Context, id int) (fruitApi.DataFruit, error)
	InsertData(ctx context.Context, data fruitApi.DataFruit) error
	DeleteDataByID(ctx context.Context, id int) error
}
