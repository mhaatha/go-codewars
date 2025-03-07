package kata

func minSlice(s []int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func maxSlice(s []int) int {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func removeZeros(s []int) []int {
	result := []int{}
	for _, v := range s {
		if v != 0 {
			result = append(result, v)
		}
	}
	return result
}

func QueueTime(customers []int, n int) int {
	totalTime := 0

	if n == 1 {
		for _, customer := range customers {
			totalTime += customer
		}
	} else if n > len(customers) {
		totalTime += maxSlice(customers)
	} else {
		currentCustomer := make([]int, 0, n)

		for i, customer := range customers {
			if len(currentCustomer) != n {
				currentCustomer = append(currentCustomer, customer)
			}
			if len(currentCustomer) == n || i == len(customers)-1 {
				smallestTime := minSlice(currentCustomer)
				totalTime += smallestTime

				for index := range currentCustomer {
					currentCustomer[index] -= smallestTime
				}

				currentCustomer = removeZeros(currentCustomer)
			}
		}
		if len(currentCustomer) != 0 {
			biggestTime := maxSlice(currentCustomer)
			totalTime += biggestTime
		}
	}

	return totalTime
}
