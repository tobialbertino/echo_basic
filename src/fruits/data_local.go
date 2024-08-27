package fruits

import (
	"context"
	"echo_basic/infra/fruitApi"
)

var LocalData = make([]fruitApi.DataFruit, 0)

func init() {
	fruitApiSvc := fruitApi.NewService()
	res, err := fruitApiSvc.GetAllData(context.Background())
	if err != nil {
		println(err.Error())
	}

	for _, val := range res {
		LocalData = append(LocalData, val)
	}
}
