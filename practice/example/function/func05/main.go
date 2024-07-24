package main

import "fmt"

type T func(int) int

func transformNumber(numbers *[]int, transform T) []int {
	res := make([]int, 0)
	for _, val := range *numbers {
		res = append(res, transform(val))
	}
	return res
}

func doubleNumbers(number int) int {
	return number * 2
}

func tripleNumbers(number int) int {
	return number * 3
}

func main() {
	input := []int{1, 2, 3, 4}
	double := transformNumber(&input, doubleNumbers)
	triple := transformNumber(&input, tripleNumbers)
	fmt.Println(double)
	fmt.Println(triple)
}
