package kata

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
		Name       string
		BaseNumber int
		Times      int
		Want       []int
	}{
		{"first test", 1, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"second test", 2, 5, []int{2, 4, 6, 8, 10}},
		{"third test", 100, 5, []int{100, 200, 300, 400, 500}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := CountBy(c.BaseNumber, c.Times)
			assertEqual(t, c.Name, got, c.Want)
		})
	}
}

func assertEqual(t testing.TB, name string, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("case %q failed: got %v, want %v", name, got, want)
	}
}
