package iteration

func Repeat(character string, numberOfIteration int) string {
	var repeated string
	for i := 0; i < numberOfIteration; i++ {
		repeated += character
	}
	return repeated
}
