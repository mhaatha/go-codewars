package kata

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}
	for _, squared := range array1 {
		currentSquared := squared * squared
		for j, squareOf := range array2 {
			if currentSquared == squareOf {
				array2 = remove(array2, j)
				break
			}
		}
	}

	return len(array2) == 0
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

/*
a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
[121 14641 20736 361 25921 361 20736 361]
[121 361 20736 361 25921 361 20736 361]

Comp(a, b) returns true because in b 121 is the square of 11, 14641 is the square of 121, and so on
It gets obvious if we write b's elements in terms of squares:
a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]
*/
