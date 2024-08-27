package words

type WordEntity struct {
	Word         string `json:"word"`
	Length       int    `json:"length"`
	NumOfVocals  int    `json:"num_of_vocals"`
	IsPalindrome bool   `json:"palindrome"`
}
