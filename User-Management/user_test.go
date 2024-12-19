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

func assertCheckId(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got ID %q want ID %q", got, want)
	}
}
