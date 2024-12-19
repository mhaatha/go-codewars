package usermanagement

type User struct {
	ID    string
	Name  string
	Email string
}

var user []User

func AddUser(u User) (string, error) {
	user = append(user, u)

	return u.ID, nil
}
