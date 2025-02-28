package kata

import "testing"

func TestCentury(t *testing.T) {
	cases := []struct {
		Name string
		Year string
		Want string
	}{
		{
			Name: "it should return 20th",
			Year: "1999",
			Want: "20th",
		},
		{
			Name: "it should return 21st",
			Year: "2011",
			Want: "21st",
		},
		{
			Name: "it should return 22nd",
			Year: "2154",
			Want: "22nd",
		},
		{
			Name: "it should return 23rd",
			Year: "2259",
			Want: "23rd",
		},
		{
			Name: "it should return 12th",
			Year: "1124",
			Want: "12th",
		},
		{
			Name: "it should return 20th",
			Year: "2000",
			Want: "20th",
		},
		{
			Name: "it should return 52nd",
			Year: "5200",
			Want: "52nd",
		},
		{
			Name: "it should return 11th",
			Year: "1100",
			Want: "11th",
		},
		{
			Name: "it should return 21st",
			Year: "2100",
			Want: "21st",
		},
		{
			Name: "it should return 12th",
			Year: "1200",
			Want: "12th",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := WhatCentury(c.Year)

			if got != c.Want {
				t.Errorf("got %v want %v", got, c.Want)
			}
		})
	}
}
