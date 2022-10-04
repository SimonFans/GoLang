package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string to bool
	var str string = "true"
	b, _ := strconv.ParseBool(str)
	fmt.Printf("type=%T, value=%v\n", b, b)
	// string to int
	var str2 string = "1001"
	num, _ := strconv.ParseInt(str2, 10, 64)
	num2 := int(num)
	num3, _ := strconv.ParseInt(str2, 2, 64)
	fmt.Printf("type=%T, value=%v\n", num, num)
	fmt.Printf("type=%T, value=%v\n", num2, num2)
	fmt.Printf("type=%T, value=%v\n", num3, num3)
	// string to float
	var str3 string = "123.456"
	f1, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("type=%T, value=%v\n", f1, f1)
}
