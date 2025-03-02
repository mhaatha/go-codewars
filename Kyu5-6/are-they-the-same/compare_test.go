package kata

import "testing"

func TestComp(t *testing.T) {
	cases := []struct {
		Name        string
		FirstArray  []int
		SecondArray []int
		Want        bool
	}{
		{
			Name:        "first test should return true",
			FirstArray:  []int{121, 144, 19, 161, 19, 144, 19, 11},
			SecondArray: []int{121, 14641, 20736, 361, 25921, 361, 20736, 361},
			Want:        true,
		},
		{
			Name:        "second test should return false",
			FirstArray:  []int{121, 144, 19, 161, 19, 144, 19, 11},
			SecondArray: []int{132, 14641, 20736, 361, 25921, 361, 20736, 361},
			Want:        false,
		},
		{
			Name:        "third test should return false",
			FirstArray:  []int{121, 144, 19, 161, 19, 144, 19, 11},
			SecondArray: []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			Want:        false,
		},
		{
			Name:        "fourth test should return false",
			FirstArray:  nil,
			SecondArray: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			Want:        false,
		},
		{
			Name:        "fifth test should return true",
			FirstArray:  []int{0, -14, 191, 161, 19, 144, 195, 1},
			SecondArray: []int{1, 0, 14 * 14, 191 * 191, 161 * 161, 19 * 19, 144 * 144, 195 * 195},
			Want:        true,
		},
		{
			Name:        "sixth test should return false",
			FirstArray:  []int{121, 144, 19, 161, 19, 144, 19, 11},
			SecondArray: []int{132, 14641, 20736, 361, 25921, 361, 20736, 361},
			Want:        false,
		},
		{
			Name:        "seventh test should return false",
			FirstArray:  []int{121, 144, 19, 161, 19, 144, 19, 11},
			SecondArray: []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361},
			Want:        false,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Comp(c.FirstArray, c.SecondArray)

			if got != c.Want {
				t.Errorf("got %v want %v", got, c.Want)
			}
		})
	}
}
