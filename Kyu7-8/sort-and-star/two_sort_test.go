package sortandstar_test

import (
	"testing"

	kata "github.com/mhaatha/go-codewars/Kyu7-8/sort-and-star"
)

func TestTwoSort(t *testing.T) {
	cases := []struct {
		Name string
		List []string
		Want string
	}{
		{"first test", []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}, "b***i***t***c***o***i***n"},
		{"second test", []string{"lets", "talk", "about", "javascript", "the", "best", "language"}, "a***b***o***u***t"},
		{"third test", []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}, "L***e***t***s"},
		{"fourth test", []string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}, "c***o***d***e"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := kata.TwoSort(c.List)

			if got != c.Want {
				t.Errorf("got %s, want %s", got, c.Want)
			}
		})
	}
}
