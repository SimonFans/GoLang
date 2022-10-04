package main

import "fmt"

func main() {

	i := 5
	fmt.Printf("The address of i=%v\n", &i)
	ptr := &i
	fmt.Printf("The ptr=%v\n", ptr)
	fmt.Printf("The address of ptr=%v\n", &ptr)
	fmt.Printf("ptr points to the value=%v\n", *ptr)
	// update the i's value
	*ptr = 10
	fmt.Printf("The updated i's value=%v\n", i)
	// 值类型: int, float, string, bool, array, struct
	// 引用类型: pointer, slice, map, chan, interface
}
