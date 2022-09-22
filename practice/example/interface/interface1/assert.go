package main

import "fmt"

func Accept(items ...int) {
	fmt.Println(items)
	fmt.Printf("items type %T", items)
}

func main() {
	var x interface{}
	var val1 float64 = 2.0
	x = val1
	// duan yan (assert)
	if y, ok := x.(float64); ok {
		fmt.Printf("y type is %T, y value is %v\n", y, y)
	} else {
		fmt.Println("fail")
	}
	fmt.Println("code continues to run...")
	Accept(8, 2, 3, 5)
}
