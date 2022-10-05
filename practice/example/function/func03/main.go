package main

import "fmt"

// define a global anonymous function
var (
	Func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

// anonymous function
func main() {
	// way 1 to call anonymous function
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Printf("n1+n2=%v\n", res1)
	// way 2 to call anonymous function
	myFun := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Printf("n1+n2=%v\n", myFun(20, 80))
	// way 3 to call global anonymous function
	fmt.Printf("n1*n2=%v\n", Func1(20, 80))
}
