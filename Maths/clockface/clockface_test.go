package clockface_test

import (
	"testing"
	"time"

	"github.com/mhaatha/go-codewars/Maths/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
	// tm's output: 1337-01-01 00:00:00 +0000 UTC
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
