package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(nums))
}

/* -------------------------------------------------------*/
func Sum(nums []int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func Sum2(nums [5]int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

/* -------------------------------------------------------*/
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sum := make([]int, lengthOfNumbers)

	for i,numbers := range numbersToSum {
		sum[i] = Sum(numbers)
 	}
	return sum
}

// REFACTOR

func SumAll2(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
