package kata_test

import (
	"reflect"
	"testing"

	kata "github.com/mhaatha/go-codewars/Kyu7-8/sort-numbers"
)

func TestSortNumbers(t *testing.T) {
	cases := []struct {
		Name          string
		Numbers, Want []int
	}{
		{"it should return a sorted numbers of slice", []int{1, 2, 10, 50, 5}, []int{1, 2, 5, 10, 50}},
		{"it should return an empty slice", []int{}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := kata.SortNumbers(c.Numbers)

			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("got %v, want %v", got, c.Want)
			}
		})
	}
}
