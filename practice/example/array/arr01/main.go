package main

import "fmt"

// value transmit
func test(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {
	(*arr)[0] = 99
}

func main() {
	// declare an array
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3, 5, 8}
	arr3 := [...]string{"a", "b"}
	arr4 := [...]string{1: "a", 0: "c", 2: "b"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	// for...range iterate an array
	for i, v := range arr1 {
		fmt.Printf("Index=%v, value=%v\n", i, v)
	}
	// value copy. Not change the value in the original array
	test(arr1)
	fmt.Printf("array value transmit=%v\n", arr1)
	// want to update the value, use pointer
	test02(&arr1)
	fmt.Printf("array pointer transmit=%v\n", arr1)
	// Print 26 alphabets
	var mychars [26]byte
	for i := 0; i < 26; i++ {
		mychars[i] = 'A' + byte(i) // 'A' + 0 byte = 65, 'A' + 1 byte = 66
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c,", mychars[i])
	}
	fmt.Println()
}
