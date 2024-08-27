package users

import (
	"context"
	"github.com/labstack/echo/v4"
)

type IDelivery interface {
	GetListData(c echo.Context) error
	GetByID(c echo.Context) error
	InsertData(c echo.Context) error
	UpdateDataByID(c echo.Context) error
	DeleteDataByID(c echo.Context) error
}

type IService interface {
	GetListData(ctx context.Context) []UserEntity
	GetByID(ctx context.Context, id int) (UserEntity, error)
	InsertData(ctx context.Context, user UserEntity) error
	UpdateDataByID(ctx context.Context, id int, user UserEntity) error
	DeleteDataByID(ctx context.Context, id int) error
}
