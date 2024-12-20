package usermanagement

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	t.Run("non-existing user", func(t *testing.T) {
		manager := NewUserManager()
		user := User{ID: "123", Name: "Test", Email: "test@example.com"}
		got, err := manager.AddUser(user)
		assertCheckId(t, got, "123")

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("existing user", func(t *testing.T) {
		manager := NewUserManager()
		manager.AddUser(User{ID: "123", Name: "Test", Email: "test@example.com"})
		got, err := manager.AddUser(User{ID: "123", Name: "Test", Email: "test@example.com"})

		if got != "" {
			t.Errorf("got %v want %v", got, "")
		}

		if err == nil {
			t.Error("expected error but success")
		}
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("users exist", func(t *testing.T) {
		manager := NewUserManager()
		user1 := User{ID: "123", Name: "Test-1", Email: "test1@example.com"}
		user2 := User{ID: "321", Name: "Test-2", Email: "test2@example.com"}
		manager.AddUser(user1)
		manager.AddUser(user2)

		got, err := manager.GetAllUser()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(got) != 2 {
			t.Errorf("expected 2 data but only %v", len(got))
		}

		if got[0] != user1 {
			if got[1] != user2 {
				t.Errorf("got %v want [%v, %v]", got, user1, user2)
			}
			t.Errorf("got %v want %v", got[0], user1)
		}
	})

	t.Run("users don't exist", func(t *testing.T) {
		manager := NewUserManager()

		got, _ := manager.GetAllUser()

		if len(got) != 0 {
			t.Errorf("got %v want %v", got, []User{})
		}
	})
}

func TestGetUserById(t *testing.T) {
	t.Run("user exists", func(t *testing.T) {
		manager := NewUserManager()
		manager.AddUser(User{ID: "123", Name: "Test", Email: "test@example.com"})

		user, err := manager.GetUserById("123")

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if user.ID != "123" {
			t.Errorf("got %v want %v", user.ID, "123")
		}
	})

	t.Run("user doesn't exists", func(t *testing.T) {
		manager := NewUserManager()

		user, err := manager.GetUserById("123")

		if user != nil {
			t.Errorf("got %v want %v", user, nil)
		}

		if err.Error() != "ID not found" {
			t.Error("expected error ID not found but found")
		}
	})
}

func TestDeleteUserById(t *testing.T) {
	t.Run("user exists", func(t *testing.T) {
		manager := NewUserManager()
		user := User{ID: "123", Name: "Test-1", Email: "test1@example.com"}
		manager.AddUser(user)

		err := manager.DeleteUserById("123")

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("user doesn't exists", func(t *testing.T) {
		manager := NewUserManager()

		err := manager.DeleteUserById("123")

		if err == nil {
			t.Errorf("expected error: %v", "ID not found")
		}
	})
}

func assertCheckId(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got ID %q want ID %q", got, want)
	}
}
