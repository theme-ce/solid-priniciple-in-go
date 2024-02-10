package main

import "fmt"

func sumOfFilteredNumbers(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number > 10 {
			sum += number
		}
	}
	return sum
}

func filterNumber(numbers []int, filter func(int) bool) []int {
	var filteredNumbers []int
	for _, number := range numbers {
		if filter(number) {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return filteredNumbers
}

func sumOfNumbers(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	sum := sumOfFilteredNumbers(numbers)
	fmt.Println(sum)

	fileterGt10Fn := func(number int) bool {
		return number > 10
	}
	filteredNumbers := filterNumber(numbers, fileterGt10Fn)

	sum = sumOfNumbers(filteredNumbers)
	fmt.Println(sum)

	fileterLt10Fn := func(number int) bool {
		return number < 10
	}
	filteredNumbers = filterNumber(numbers, fileterLt10Fn)

	sum = sumOfNumbers(filteredNumbers)
	fmt.Println(sum)
}
