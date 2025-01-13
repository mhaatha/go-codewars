package main

import "testing"

func BenchmarkArabicToRomans(b *testing.B) {
	cases := []int{1, 100, 3999}
	for _, c := range cases {
		b.Run("arabics to romans", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConvertToRoman(c)
			}
		})
	}
}

func BenchmarkRomansToArabic(b *testing.B) {
	cases := []string{"I", "C", "MMMCMXCIX"}
	for _, c := range cases {
		b.Run("romans to arabics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConvertToArabic(c)
			}
		})
	}
}
