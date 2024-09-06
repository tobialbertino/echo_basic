package products

import (
	"context"
	"echo_basic/infra/gemini"
	"github.com/labstack/echo/v4"
)

type IDelivery interface {
	GetListData(c echo.Context) error
	GetByID(c echo.Context) error
	InsertData(c echo.Context) error
	UpdateDataByID(c echo.Context) error
	DeleteDataByID(c echo.Context) error

	// Gemini API
	GetContent(c echo.Context) error
}

type IService interface {
	GetListData(ctx context.Context) []ProductEntity
	GetByID(ctx context.Context, id string) (ProductEntity, error)
	InsertData(ctx context.Context, user ProductEntity) error
	UpdateDataByID(ctx context.Context, id string, user ProductEntity) error
	DeleteDataByID(ctx context.Context, id string) error

	// Gemini API
	GetContent(ctx context.Context, content string) (gemini.RespContent, error)
}
