package main

import "testing"

func TestGreeting(t *testing.T) {

	got := greeting()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreeting_recipient_test(t *testing.T) {

	got := greeting_recipirnt("Khalil")
	want:= "Hello, Khalil!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}