package main

/*
Example.
Find sum of numbers
*/

/*
Bad practice.
This function is not following the Liskov Substitution Principle.
It is tightly coupled with the int type.
If we want to sum a list of float64, we need to modify this function.
*/

func processNumber(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

/*
Good practice.
This function is following the Liskov Substitution Principle.
It is loosely coupled with the Number interface.
If we want to sum a list of float64, we can pass a new function to it.
*/

type Number interface {
	~int | ~float64
}

func sumNumbersGeneric[T Number](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

type ProcessFn[T Number] func(nums []T) T

func processNumbersGeneric[T Number](nums []T, processFn ProcessFn[T]) T {
	return processFn(nums)
}

func main() {
	_ = processNumber([]int{1, 2})

	_ = processNumbersGeneric([]int{1, 2}, sumNumbersGeneric)
	_ = processNumbersGeneric([]float64{1.1, 2.2}, sumNumbersGeneric)
}
