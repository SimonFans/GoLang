package main

import "fmt"

type myInt int
type funcvar func(int, int) int

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func mySumFun(f funcvar, n1 int, n2 int) int {
	return f(n1, n2)
}

// define function mySum
func getSum2(n1 int, args ...int) int {
	total := n1
	for i := 0; i < len(args); i++ {
		total += args[i]
	}
	return total
}

func main() {
	// self define data type still need to convert to basic data type when use it
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)
	fmt.Printf("num1=%v, num2=%v\n", num1, num2)
	// self define function type example
	n1, n2 := 5, 10
	_sum := mySumFun(getSum, n1, n2)
	fmt.Println(_sum)
	// Go supports variable params. Example: write a function sum, to get the sum of  1 to multiple int values
	mySum := getSum2(10, 0, -1, -2)
	fmt.Printf("mySum=%v\n", mySum)
}
