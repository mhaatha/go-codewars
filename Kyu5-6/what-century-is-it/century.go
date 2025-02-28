package kata

import (
	"strconv"
)

// WhatCentury handles year in 4 digits.
// If you pass the digit less or more than 4, that's not gonna work.
func WhatCentury(year string) string {
	firstTwo := getTheFirstTwo(year)

	// e.g: 2000 -> 20th
	// e.g: 1000 -> 10th
	if string(year[1:]) == "000" {
		return firstTwo + "th"
	}
	if string(year[2:]) == "00" {
		if string(firstTwo[0]) == "1" {
			return firstTwo + "th"
		} else if string(firstTwo[1]) != "1" {
			switch string(firstTwo[1]) {
			case "1":
				return firstTwo + "st"
			case "2":
				return firstTwo + "nd"
			case "3":
				return firstTwo + "rd"
			default:
				return firstTwo + "th"
			}
		} else {
			switch string(firstTwo[1]) {
			case "1":
				return firstTwo + "st"
			case "2":
				return firstTwo + "nd"
			case "3":
				return firstTwo + "rd"
			default:
				return firstTwo + "th"
			}
		}
	}

	// Convert the firstTwo into integer
	intYear, err := strconv.Atoi(firstTwo)
	if err != nil {
		panic(err)
	}

	intYear += 1

	// e.g: 1190 -> 12th
	// e.g: 1222 -> 13th
	if string(firstTwo[0]) == "1" {
		// Convert the intYear into string
		return strconv.Itoa(intYear) + "th"
	}

	switch string(firstTwo[1]) {
	case "0":
		return strconv.Itoa(intYear) + "st"
	case "1":
		return strconv.Itoa(intYear) + "nd"
	case "2":
		return strconv.Itoa(intYear) + "rd"
	}

	return strconv.Itoa(intYear) + "th"
}

func getTheFirstTwo(year string) string {
	return string(year[0]) + string(year[1])
}

/*
"1999" --> "20th"
"2011" --> "21st"
"2154" --> "22nd"
"2259" --> "23rd"
"1124" --> "12th"
"2000" --> "20th"

1999 -> 20

2011 -> 21

2154 -> 22

2259 -> 23

1124 -> 12

2000 -> 20
*/
