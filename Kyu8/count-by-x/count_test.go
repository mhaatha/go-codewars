package kata

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
		Name             string
		NumberToMultiply int
		Multiplier       int
		Want             []int
	}{
		{"First test", 1, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Second test", 2, 5, []int{2, 4, 6, 8, 10}},
		{"Third test", 100, 5, []int{100, 200, 300, 400, 500}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := CountBy(c.NumberToMultiply, c.Multiplier)
			assertEqual(t, got, c.Want)
		})
	}
}

func assertEqual(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
