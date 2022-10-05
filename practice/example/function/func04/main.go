package main

import (
	"fmt"
	"strconv"
)

// bi bao example
// because we use n in the anonymous function func(x int) int, so n will be aggregated everytime instead initialing everytime.
func accumulator() func(int) int {
	var n int = 10
	var str = "Hello"
	return func(x int) int {
		n = n + x
		str += strconv.Itoa(x)
		fmt.Println("str=", str)
		return n
	}
}

func main() {
	f := accumulator()
	fmt.Println(f(1))
	fmt.Println(f(3))
}
