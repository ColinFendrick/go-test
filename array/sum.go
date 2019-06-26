package main

// Sum all numbers in slice
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAllTails numbers in slices
func SumAllTails(numbersArray ...[]int) (sums []int) {
	for _, numbers := range numbersArray {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
