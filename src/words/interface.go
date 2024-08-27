package words

import (
	"context"
	"github.com/labstack/echo/v4"
)

// Domain Users
type IDelivery interface {
	GetAllData(c echo.Context) error
	InsertData(c echo.Context) error
}

type IService interface {
	GetAllData(ctx context.Context) []WordEntity
	InsertData(ctx context.Context, word string) error
}
