package utils

func IsPalindrome(word string) bool {
	lenStr := len(word)
	for i := 0; i < lenStr/2; i++ {
		if word[i] != word[lenStr-i-1] {
			return false
		}
	}
	return true
}

func LenVocals(word string) (res int) {
	var vocal map[rune]bool = map[rune]bool{
		'a': true,
		'i': true,
		'u': true,
		'e': true,
		'o': true,
	}

	for _, char := range word {
		if vocal[char] {
			res++
		}
	}
	return
}

func IsEmptyString(s string) bool {
	if s == "" && len(s) == 0 {
		return true
	}
	return false
}
