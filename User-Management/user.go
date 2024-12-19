package usermanagement

import "errors"

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

func GetUserById(id string) (string, error) {
	for i := range UserSlice {
		if UserSlice[i].ID == id {
			return UserSlice[i].ID, nil
		}
	}

	return "", errors.New("ID not found")
}

func DeleteUserById(id string) error {
	for i := range UserSlice {
		if UserSlice[i].ID == id {
			UserSlice = append(UserSlice[:i], UserSlice[i+1:]...)
			return nil
		}
	}

	return errors.New("ID not found")
}
