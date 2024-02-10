package main

/*
Example.
Find sum of a filtered list of numbers that greater than 10.
*/

/*
Bad practice.
This function is not following the Single Responsibility Principle.
It is doing two things: filtering and summing.
*/

func sumOfFilteredNumbers(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number > 10 {
			sum += number
		}
	}

	return sum
}

/*
Good practice.
This function is following the Single Responsibility Principle.
It is doing one thing: filtering.
*/

func filterNumbers(numbers []int) []int {
	var filteredNumbers []int
	for _, number := range numbers {
		if number > 10 {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return filteredNumbers
}

/*
Good practice.
This function is following the Single Responsibility Principle.
It is doing one thing: summing.
*/

func sumNumbers(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	_ = sumOfFilteredNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	filteredNumbers := filterNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	_ = sumNumbers(filteredNumbers)
}
