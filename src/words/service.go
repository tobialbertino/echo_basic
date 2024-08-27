package words

import (
	"context"
	"echo_basic/pkg/utils"
	"errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAllData(_ context.Context) []WordEntity {
	return LocalData
}

func (s *Service) InsertData(_ context.Context, word string) (err error) {
	// check empty string
	if utils.IsEmptyString(word) {
		err = errors.New(utils.MsgFailedToAddWord)
		return
	}

	newData := WordEntity{
		Word:         word,
		Length:       len(word),
		NumOfVocals:  utils.LenVocals(word),
		IsPalindrome: utils.IsPalindrome(word),
	}

	LocalData = append(LocalData, newData)

	return
}
