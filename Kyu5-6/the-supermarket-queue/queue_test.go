package kata

import "testing"

func TestQueue(t *testing.T) {
	cases := []struct {
		Name      string
		Customers []int
		Tills     int
		Want      int
	}{
		{
			Name:      "it should return 0",
			Customers: []int{},
			Tills:     1,
			Want:      0,
		},
		{
			Name:      "it should return 10",
			Customers: []int{1, 2, 3, 4},
			Tills:     1,
			Want:      10,
		},
		{
			Name:      "it should return 9",
			Customers: []int{2, 2, 3, 3, 4, 4},
			Tills:     2,
			Want:      9,
		},
		{
			Name:      "it should return 12",
			Customers: []int{2, 3, 10},
			Tills:     2,
			Want:      12,
		},
		{
			Name:      "it should return 5",
			Customers: []int{1, 2, 3, 4, 5},
			Tills:     100,
			Want:      5,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := QueueTime(c.Customers, c.Tills)

			if got != c.Want {
				t.Errorf("got %v want %v", got, c.Want)
			}
		})
	}
}
