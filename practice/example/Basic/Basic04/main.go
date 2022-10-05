package main

import "fmt"

func main() {
	// 2
	fmt.Println(10 / 4)
	// 2.5
	fmt.Println(10 / 4.0)
	// 1
	fmt.Println(10 % 3)
	// -1 ==> a%b = a - a/b*b
	fmt.Println(-10 % 3)
	// swap numbers
	a, b := 2, 9
	a, b = b, a
	fmt.Println(a, b)
	// bit operator
	// << => num * 2^(), >> => num / 2^()
	num := 3
	num = num << 2
	fmt.Println(num) // 12
	num2 := 24
	num2 = num2 >> 2
	fmt.Println(num2) // 6
}
