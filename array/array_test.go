package main

import (
	"slices"
	"testing"
)

func TestSum (t *testing.T){

	t.Run("Collection of 5 numbers", func (t *testing.T) {
		numbers := [5]int{1,2,3,4,5}
        output := Sum2(numbers)
		expected := 15

		if output != expected {
			t.Errorf("Expected %d, but Output is %d, given %v", expected, output, numbers)
		}
	})

	t.Run("Collection of any size", func (t *testing.T) {
		numbers := []int {1,2,3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	
}

func TestSumAll (t *testing.T){
	got := SumAll([]int{1,2}, []int{0,9})
	want:= []int{3,9}

	if !slices.Equal(got,want) {
		t.Errorf("got %v want %v", got, want)
	}
}