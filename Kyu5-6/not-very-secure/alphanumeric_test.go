package kata

import "testing"

func TestAlphanumeric(t *testing.T) {
	cases := []struct {
		Name string
		Str  string
		Want bool
	}{
		{
			Name: "First test",
			Str:  ".*?",
			Want: false,
		},
		{
			Name: "Second test",
			Str:  "Mazinkaiser",
			Want: true,
		},
		{
			Name: "Third test",
			Str:  "hello world_",
			Want: false,
		},
		{
			Name: "Fourth test",
			Str:  "     ",
			Want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Alphanumeric(c.Str)
			if got != c.Want {
				t.Errorf("got '%v' want '%v'", got, c.Want)
			}
		})
	}
}
