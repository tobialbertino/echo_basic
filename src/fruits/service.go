package fruits

import (
	"context"
	"echo_basic/infra/fruitApi"
	"errors"
)

type Service struct {
}

func NewService() IService {
	return &Service{}
}

func (s *Service) GetAllData(_ context.Context) ([]fruitApi.DataFruit, error) {
	return LocalData, nil
}

func (s *Service) GetDataByID(_ context.Context, id int) (res fruitApi.DataFruit, err error) {
	for _, data := range LocalData {
		if data.ID == id {
			res = data
			return
		}
	}

	err = errors.New("not found")
	return
}

func (s *Service) InsertData(_ context.Context, data fruitApi.DataFruit) (err error) {
	LocalData = append(LocalData, data)
	return
}

func (s *Service) DeleteDataByID(_ context.Context, id int) (err error) {
	// validate check index len slice
	var isContain bool
	for _, data := range LocalData {
		if data.ID == id {
			isContain = true
		}
	}

	if !isContain {
		err = errors.New("not found")
		return
	}

	for idx, data := range LocalData {
		if data.ID == id {
			LocalData = RemoveSLiceFruitByID(LocalData, idx)
			return
		}
	}

	err = errors.New("not found")
	return
}
