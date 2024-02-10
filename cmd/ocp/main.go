package main

/*
Example.
Filter the list of numbers on some condition.
*/

/*
Bad practice.
This function is not following the Open/Closed Principle.
If we want to filter numbers that are greater than 20, we need to modify this function.
*/

func filterNumbers(numbers []int) []int {
	var result []int
	for _, number := range numbers {
		if number > 10 {
			result = append(result, number)
		}
	}
	return result
}

/*
Good practice.
This function is following the Open/Closed Principle.
It is open for extension and closed for modification.
If we want to filter numbers that are greater than 20, we can pass a new function to it.
*/

type FilterFn func(int) bool

func filterNumbersWithFn(numbers []int, fn FilterFn) []int {
	var result []int
	for _, number := range numbers {
		if fn(number) {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	_ = filterNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	_ = filterNumbersWithFn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(n int) bool {
		return n > 10
	})
}
