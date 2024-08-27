package users

import (
	"context"
	"errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetListData(_ context.Context) []UserEntity {
	return LocalData
}

func (s *Service) GetByID(_ context.Context, id int) (res UserEntity, err error) {
	for _, data := range LocalData {
		if data.ID == id {
			res = data
			return
		}
	}

	err = errors.New("not found")
	return
}

func (s *Service) InsertData(_ context.Context, user UserEntity) (err error) {
	LocalData = append(LocalData, user)
	return
}

func (s *Service) UpdateDataByID(_ context.Context, id int, user UserEntity) (err error) {
	for i, data := range LocalData {
		if data.ID == id {
			LocalData[i] = user
			return
		}
	}

	err = errors.New("not found")
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
			LocalData = RemoveSLiceUserByID(LocalData, idx)
			return
		}
	}

	err = errors.New("not found")
	return
}
