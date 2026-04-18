package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {

	sum := Add(2,2)
	expected:= 2+2

	if sum != expected {
		t.Errorf("Expected %d, but Got %d", expected, sum)
	}
}

func ExampleAdd() {
	fmt.Println(Add(1, 5))
	// Output: 6
}