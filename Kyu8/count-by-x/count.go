package kata

func CountBy(x, n int) []int {
	result := make([]int, n)

	for i := 1; i <= n; i++ {
		result[i-1] = x * i
	}

	return result
}
