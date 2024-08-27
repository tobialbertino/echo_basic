package users

func RemoveSLiceUserByID(s []UserEntity, idx int) []UserEntity {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}
