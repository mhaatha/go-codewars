// https://www.codewars.com/kata/5648b12ce68d9daa6b000099/solutions/go

package kata

import "testing"

func TestNumber(t *testing.T) {
	cases := []struct {
		Name string
		List [][2]int
		Want int
	}{
		{"first test", [][2]int{{10, 0}, {3, 5}, {5, 8}}, 5},
		{"second test", [][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}, 17},
		{"third test", [][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}, 21},
	}

	for _, c := range cases {
		got := Number(c.List)

		if got != c.Want {
			t.Errorf("got %v want %v", got, c.Want)
		}
	}
}
