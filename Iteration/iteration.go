package iteration

func Repeat(letter string, loopTimes int) string {
	var repeated string

	for i := 0; i < loopTimes; i++ {
		repeated += letter
	}
	return repeated
}
