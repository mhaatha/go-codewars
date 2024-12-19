package usermanagement

import (
	"testing"

	"github.com/google/uuid"
)

func TestAddUser(t *testing.T) {
	newUUID := uuid.NewString()
	user := User{
		ID:    newUUID,
		Name:  "Test 1",
		Email: "test1@example.com",
	}

	want := newUUID
	got, err := AddUser(user)

	assertCheckId(t, got, want)

	if err != nil {
		t.Errorf("got error %q", err)
	}
}

func TestGetAllUsers(t *testing.T) {
	newUUID := uuid.NewString()
	user := User{
		ID:    newUUID,
		Name:  "Test 1",
		Email: "test1@example.com",
	}
	UserSlice = append(UserSlice, user)

	got, err := GetAllUser()
	want := UserSlice

	assertCheckUsers(t, got, want)

	if err != nil {
		t.Errorf("got error %q", err)
	}
}

func TestGetUserById(t *testing.T) {
	newUUID := uuid.NewString()
	user := User{
		ID:    newUUID,
		Name:  "Test 1",
		Email: "test1@example.com",
	}
	UserSlice = append(UserSlice, user)

	got, err := GetUserById(newUUID)
	want := newUUID

	assertCheckId(t, got, want)

	if err != nil {
		t.Errorf("got error %q", err)
	}
}

func TestDeleteUserById(t *testing.T) {
	newUUID := uuid.NewString()
	user := User{
		ID:    newUUID,
		Name:  "Test 1",
		Email: "test1@example.com",
	}
	UserSlice = append(UserSlice, user)

	err := DeleteUserById(newUUID)

	if err != nil {
		t.Errorf("got error %q", err)
	}
}

func assertCheckId(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got ID %q want ID %q", got, want)
	}
}

func assertCheckUsers(t testing.TB, got, want []User) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %q want %q", got, want)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
