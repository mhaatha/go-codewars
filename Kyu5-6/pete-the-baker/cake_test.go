package kata

import (
	"testing"
)

func TestCake(t *testing.T) {
	cases := []struct {
		Name              string
		Recipe, Available map[string]int
		Want              int
	}{
		{
			Name:      "First test",
			Recipe:    map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
			Available: map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200},
			Want:      2,
		},
		{
			Name:      "Second test",
			Recipe:    map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100},
			Available: map[string]int{"sugar": 500, "flour": 2000, "milk": 2000},
			Want:      0,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Cake(c.Recipe, c.Available)

			if c.Want != got {
				t.Errorf("got %v want %v", got, c.Want)
			}
		})
	}
}
