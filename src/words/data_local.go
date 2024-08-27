package words

var LocalData = make([]WordEntity, 0)

func init() {
	LocalData = append(LocalData, WordEntity{
		Word:         "test",
		Length:       4,
		NumOfVocals:  1,
		IsPalindrome: false,
	})
}
