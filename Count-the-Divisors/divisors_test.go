package main

import "testing"

func TestDivisors(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		got := Divisors(1)
		want := 1

		assertCorrectCount(t, got, want)
	})
	t.Run("11", func(t *testing.T) {
		got := Divisors(11)
		want := 2

		assertCorrectCount(t, got, want)
	})
	t.Run("54", func(t *testing.T) {
		got := Divisors(54)
		want := 8

		assertCorrectCount(t, got, want)
	})
	t.Run("64", func(t *testing.T) {
		got := Divisors(64)
		want := 7

		assertCorrectCount(t, got, want)
	})
}

func assertCorrectCount(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
