# go-repo

### Discipline

Let's go over the cycle again
    - Write a test
    - Make the compiler pass
    - Run the test, see that it fails and check the error message is meaningful
    - Write enough code to make the test pass
    - Refactor

### -----------------------------------------------------------
To get the test coverage for the package, we run the test with coverage enabled by providing the -cover flag to go test:

% go test -cover

### -----------------------------------------------------------
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

You'll see the code is very similar to a test.

The testing.B gives you access to the loop function. Loop() returns true as long as the benchmark should continue running.

When the benchmark code is executed, it measures how long it takes. After Loop() returns false, b.N contains the total number of iterations that ran.

The number of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")
### -----------------------------------------------------------

#### Concatenation

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