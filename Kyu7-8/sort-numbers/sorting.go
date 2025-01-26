package kata

func SortNumbers(numbers []int) []int {
	if len(numbers) < 1 {
		return []int{}
	}

	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				swap := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = swap
			}
		}
	}

	return numbers
}
