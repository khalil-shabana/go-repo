package iteration

import (
	"testing"
)


func TestRepeat(t *testing.T){
	repeated := Repeat("a",4)
	expected := "aaaa"

	assertCorrectMessage(t, repeated,expected) 
}


// The testing.B gives you access to the loop function.
// Loop() returns true as long as the benchmark should continue running.
// When the benchmark code is executed, it measures how long it takes.
//  After Loop() returns false, b.N contains the total number of iterations that ran.

// to run : go test -bench="."

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 4)
		//fmt.Printf("%d", b.N)
	}
}
func BenchmarkRepeat2(b *testing.B) {
	for b.Loop() {
		RepeatWizStringBuilder("a")
		//fmt.Printf("%d", b.N)
	}
}




func assertCorrectMessage (t testing.TB, output, expected string){
	t.Helper()
	if output != expected {
		t.Errorf("Expected %s, but Output is %s", expected, output)
	}
}