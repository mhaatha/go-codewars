package kata

import (
	"regexp"
)

func Alphanumeric(str string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", str)

	return match
}
