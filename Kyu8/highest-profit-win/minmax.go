// https://www.codewars.com/kata/559590633066759614000063/solutions/go

package kata

func MinMax(arr []int) [2]int {
	min, max := arr[0], arr[0]

	for _, index := range arr {
		if min > index {
			min = index
		}
		if max < index {
			max = index
		}
	}

	return [2]int{min, max}
}
