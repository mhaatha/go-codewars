package kata

import "testing"

func TestLoveFunc(t *testing.T) {
	cases := []struct {
		name    string
		flower1 int
		flower2 int
		want    bool
	}{
		{"first test", 1, 4, true},
		{"second test", 2, 2, false},
		{"third test", 10, 12, false},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			got := LoveFunc(cc.flower1, cc.flower2)

			if got != cc.want {
				t.Errorf("got %v want %v", got, cc.want)
			}
		})
	}
}
