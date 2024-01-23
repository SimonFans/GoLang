package main

import "fmt"

func Substract(x int) (num int) {
	defer func(x int) {
		num = x - 10
	}(x)
	fmt.Println("- 10")
	return num
}

func test(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
	}()
	fmt.Println("x / y = ", z)
}

func main() {
	num := 20
	num = Substract(num)
	fmt.Println(num)

	test(6, 4)
	test(0, 1)
	test(1, 0)

}
