// https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a/train/go

package main

import "fmt"

func Past(h, m, s int) int {
	// Convert hour into milisecond
	h = h * 3600000
	// Convert minute into milisecond
	m = m * 60000
	// Convert second into milisecond
	s = s * 1000
	// Sum all of the hour, minute and second
	return h + m + s
}

func main() {
	fmt.Println(Past(0, 1, 1))
}
