package main

// Sum all numbers in slice
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll numbers in slices
func SumAll(numbersArray ...[]int) (sums []int) {
	for _, numbers := range numbersArray {
		sums = append(sums, Sum(numbers))
	}
	return
}
