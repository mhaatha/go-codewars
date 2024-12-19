package usermanagement

type User struct {
	ID    string
	Name  string
	Email string
}

var UserSlice []User

func AddUser(u User) (string, error) {
	UserSlice = append(UserSlice, u)

	return u.ID, nil
}

func GetAllUser() ([]User, error) {
	return UserSlice, nil
}
