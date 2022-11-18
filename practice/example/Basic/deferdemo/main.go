package main

import "fmt"

func TwoSum(n1, n2 int) int {
	defer fmt.Printf("n1=%v\n", n1) // 3
	defer fmt.Printf("n2=%v\n", n2) // 2
	res := n1 + n2
	fmt.Printf("res=%v\n", res) // 1
	return res
}

func deferTest(n int) int {
	defer fmt.Printf("defer n=%v\n", n)
	fmt.Println("start...")
	return n
}

func main() {
	res := TwoSum(10, 20)
	fmt.Printf("Main res=%v\n", res) // 4
	fmt.Println(deferTest(5))
}
