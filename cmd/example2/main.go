package main

import (
	"fmt"
	"sort"
)

type Comparator[T any] func(a, b T) bool

func Sort[T any](slice []T, comparator Comparator[T]) {
	sort.Slice(slice, func(i, j int) bool {
		return comparator(slice[i], slice[j])
	})
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	intComparator := func(a, b int) bool { return a < b }
	Sort(intSlice, intComparator)
	fmt.Println("Sorted integers:", intSlice)

	stringSlice := []string{"banana", "apple", "cherry"}
	stringComparator := func(a, b string) bool { return a < b }
	Sort(stringSlice, stringComparator)
	fmt.Println("Sorted strings:", stringSlice)
}
