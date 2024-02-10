package main

/*
Example.
Process a numbers list with a function.
*/

/*
Bad practice.
This function is not following the Dependency Inversion Principle.
It is tightly coupled with the sumNumbers function.
*/

func processNumbers(numbers []int) int {
	return sumNumbers(numbers)
}

func sumNumbers(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
Good practice.
This function is following the Dependency Inversion Principle.
It is loosely coupled with the sumNumbers function.
*/

type ProcessFn func([]int) int

func processNumbersWithFn(numbers []int, fn ProcessFn) int {
	return fn(numbers)
}

func main() {
	_ = processNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	_ = processNumbersWithFn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, sumNumbers)
}
