package kata

import (
	"strconv"
)

func NbDig(n, d int) int {
	squaredNumbers := []int{}
	counter := 0

	for i := 0; i <= n; i++ {
		squaredNumbers = append(squaredNumbers, i*i)
	}

	for i := 0; i < len(squaredNumbers); i++ {
		intToStr := strconv.Itoa((squaredNumbers[i]))
		length := len(intToStr)
		for j := 0; j < length; j++ {
			if string(intToStr[j]) == string(strconv.Itoa(d)) {
				counter++
			}
		}
	}
	return counter
}
