package products

func RemoveSLiceUserByID(s []ProductEntity, idx int) []ProductEntity {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}
