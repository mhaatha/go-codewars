package kata

func CountBy(x, n int) []int {
	result := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		result = append(result, x*i)
	}

	return result
}
