package hello

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {

	got := greeting()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Examplegreeting() {

	fmt.Println(greeting())
	// Output: Hello, World!
}
func TestGreeting_recipient(t *testing.T) {

	got := greeting_recipirnt("Khalil")
	want:= "Hello, Khalil!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T){
	t.Run("hello to a person", func (t *testing.T){
		got := Hello("Khalil", "")
		want:= "Hello, Khalil"

		if got != want {
			t.Errorf("got %q, want %q", got,want)
		}
	})

	t.Run("empty string, saying Hello, World", func (t *testing.T) {
		got := Hello("","")
		want:= "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Khalil", "Spanish")
		want:= "Hola, Khalil"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Frensh", func(t *testing.T) {
		got := Hello("Khalil", "French")
		want:= "Bonjour, Khalil"

		assertCorrectMessage(t, got, want)
	})
}





func assertCorrectMessage(t testing.TB, got string, want string){
	t.Helper()
	if got != want {
			t.Errorf("got %q, want %q", got,want)
	}
}
