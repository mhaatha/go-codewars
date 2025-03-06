package kata

import "testing"

func TestFormatDuration(t *testing.T) {
	cases := []struct {
		Name    string
		Seconds int64
		Want    string
	}{
		{
			Name:    "it should return 'now' if seconds is 0",
			Seconds: 0,
			Want:    "now",
		},
		{
			Name:    "it should return '1 minute and 2 seconds' if seconds is 62",
			Seconds: 62,
			Want:    "1 minute and 2 seconds",
		},
		{
			Name:    "it should return '1 hour, 1 minute and 2 seconds' if seconds is 3662",
			Seconds: 3662,
			Want:    "1 hour, 1 minute and 2 seconds",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := FormatDuration(c.Seconds)
			if got != c.Want {
				t.Errorf("got '%v' want '%v'", got, c.Want)
			}
		})
	}
}
