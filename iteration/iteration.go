package iteration

import "strings"

const repeatCount = 3

func Repeat(x string, r int) string {
	var repeated string
	for i := 0; i < r; i++ {
		repeated += x
	}
	return repeated
}


// Strings in Go are immutable, meaning every concatenation, such as in our Repeat function,
// involves copying memory to accommodate the new string.
// This impacts performance, particularly during heavy string concatenation.
// The standard library provides the strings.BuilderstringsBuilder type
// which minimizes memory copying.
// It implements a WriteString method which we can use to concatenate strings
func RepeatWizStringBuilder(c string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(c)
	}
	return repeated.String()
}