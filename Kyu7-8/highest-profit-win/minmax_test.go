// https://www.codewars.com/kata/559590633066759614000063/solutions/go

package kata

import (
	"reflect"
	"testing"
)

func TestMinMax(t *testing.T) {
	cases := []struct {
		Name string
		List []int
		Want [2]int
	}{
		{"first test", []int{1, 2, 3, 4, 5}, [2]int{1, 5}},
		{"second test", []int{5, 2334454}, [2]int{5, 2334454}},
		{"third test", []int{1}, [2]int{1, 1}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := MinMax(c.List)

			if !reflect.DeepEqual(got, c.Want) {
				t.Errorf("got %v, want %v", got, c.Want)
			}
		})
	}
}
