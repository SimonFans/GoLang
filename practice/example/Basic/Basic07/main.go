package main

import "fmt"

// Print 2,2,3
func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Printf("n=%v\n", n)
}

// Print 2 because test() is in the if clause
func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Printf("n=%v\n", n)
	}
}

func main() {
	test1(4)
	test2(4)
}
