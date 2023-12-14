package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	numbers := []int{12, 6, 8, 10, 55}
	slices.Sort(numbers)
	fmt.Println(numbers)
	fmt.Println(slices.IsSorted(numbers))
	myCompare := func(a, b int) int {
		return -1 * cmp.Compare(a, b)
	}
	slices.SortFunc(numbers, myCompare)
	fmt.Println(numbers)
}
