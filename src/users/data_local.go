package users

var LocalData = make([]UserEntity, 0)

func init() {
	LocalData = append(LocalData, UserEntity{
		ID:       1,
		Name:     "nama1",
		Username: "username1",
		Password: "pass1",
		Address:  "address1",
	})
}
