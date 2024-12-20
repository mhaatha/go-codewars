package usermanagement

import "errors"

type User struct {
	ID    string
	Name  string
	Email string
}

type UserManager struct {
	users []User
}

func NewUserManager() *UserManager {
	return &UserManager{users: []User{}}
}

func (um *UserManager) AddUser(u User) (string, error) {
	for _, user := range um.users {
		if user.ID == u.ID {
			return "", errors.New("ID already exists")
		}
	}
	um.users = append(um.users, u)
	return u.ID, nil
}

func (um *UserManager) GetAllUser() ([]User, error) {
	return um.users, nil
}

func (um *UserManager) GetUserById(id string) (*User, error) {
	for _, user := range um.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("ID not found")
}

func (um *UserManager) DeleteUserById(id string) error {
	for i := range um.users {
		if um.users[i].ID == id {
			um.users = append(um.users[:i], um.users[i+1:]...)
			return nil
		}
	}
	return errors.New("ID not found")
}
