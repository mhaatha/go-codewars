// https://www.codewars.com/kata/5648b12ce68d9daa6b000099/solutions/go

package kata

func Number(busStops [][2]int) int {
	peopleEnter, peopleOut := 0, 0
	for _, bs := range busStops {
		peopleEnter += bs[0]
		peopleOut += bs[1]
	}

	return peopleEnter - peopleOut
}
