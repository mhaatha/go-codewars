package main

import "testing"

func TestPast(t *testing.T) {
	got := Past(0, 1, 1)
	want := 61000

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
