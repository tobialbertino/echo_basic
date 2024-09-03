package products

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetListData(_ context.Context) []ProductEntity {
	return LocalData
}

func (s *Service) GetByID(_ context.Context, id string) (res ProductEntity, err error) {
	for _, data := range LocalData {
		if data.ID == id {
			res = data
			return
		}
	}

	err = errors.New("not found")
	return
}

func (s *Service) InsertData(_ context.Context, user ProductEntity) (err error) {
	user.ID = uuid.New().String()
	LocalData = append(LocalData, user)
	return
}

func (s *Service) UpdateDataByID(_ context.Context, id string, user ProductEntity) (err error) {
	for i, data := range LocalData {
		if data.ID == id {
			LocalData[i] = user
			return
		}
	}

	err = errors.New("not found")
	return
}

func (s *Service) DeleteDataByID(_ context.Context, id string) (err error) {
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
			LocalData = RemoveSLiceUserByID(LocalData, idx)
			return
		}
	}

	err = errors.New("not found")
	return
}
