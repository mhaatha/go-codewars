package kata

import "testing"

func TestNbDig(t *testing.T) {
	cases := []struct {
		name       string
		n, d, want int
	}{
		{"first test case", 10, 1, 4},
		{"second test case", 25, 1, 11},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			got := NbDig(cc.n, cc.d)

			if got != cc.want {
				t.Errorf("got %v, want %v", got, cc.want)
			}
		})
	}
}
