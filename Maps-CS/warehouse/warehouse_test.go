package warehouse

import "testing"

func TestWarehouse(t *testing.T) {
	wh := NewWarehouse()

	t.Run("Add new item", func(t *testing.T) {
		err := wh.AddItem("Laptop", 10)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		got, _ := wh.CheckStock("Laptop")
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Update existing item stock", func(t *testing.T) {
		wh.AddItem("Phone", 20)
		err := wh.UpdateStock("Phone", 15)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		got, _ := wh.CheckStock("Phone")
		want := 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Check non-existing item stock", func(t *testing.T) {
		_, err := wh.CheckStock("Tablet")
		if err == nil {
			t.Errorf("expected error for non-existing item, got nil")
		}
	})

	t.Run("Delete item", func(t *testing.T) {
		wh.AddItem("Headphones", 5)
		err := wh.DeleteItem("Headphones")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		_, err = wh.CheckStock("Headphones")
		if err == nil {
			t.Error("expected error fot deleted item, got nil")
		}
	})
}
