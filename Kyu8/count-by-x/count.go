package kata

func CountBy(x, n int) []int {
	result := []int{}
	numberToPush := x

	for i := 0; i < n; i++ {
		result = append(result, numberToPush)
		numberToPush += x
	}

	return result
}
