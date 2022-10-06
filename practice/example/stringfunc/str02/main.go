package main

import "fmt"

func main() {
	// modify string value
	// 1. string -> []byte or []rune -> update value -> convert into string()
	o_str1 := "abcd"
	arr := []byte(o_str1)
	arr[0] = 'z'
	n_str1 := string(arr)
	fmt.Printf("Old string=%v, New string=%v\n", o_str1, n_str1)

	// 2. deal with Chinese characters. Because one Chinese characters take three bytes, use []byte will generate unidentified characters
	o_str2 := "abcd"
	arr2 := []rune(o_str2)
	arr2[0] = 'z'
	n_str2 := string(arr2)
	fmt.Printf("Old string=%v, New string=%v\n", o_str2, n_str2)
}
