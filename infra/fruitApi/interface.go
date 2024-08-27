package fruitApi

import "context"

type IService interface {
	GetAllData(ctx context.Context) (res []DataFruit, err error)
}
