package main

import "fmt"

func test(char byte) byte {
	return char + 1
}

func main() {

	var key byte
	fmt.Println("Enter one character. a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	fmt.Printf("return type=%T, value=%v\n", test(key), test(key))
	switch test(key) {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	default:
		fmt.Println("Default")
	}
	// switch fallthrough example
	/*
		match 10
		match 20
		match 30
	*/
	num := 10
	switch num {
	case 10:
		fmt.Println("match 10")
		fallthrough
	case 20:
		fmt.Println("match 20")
		fallthrough
	case 30:
		fmt.Println("match 30")
	default:
		fmt.Println("no match")

	}
	// Type switch can be used to determine interface(variable) points to which variable type
	var x interface{}
	y := 10
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x is type=%T", i)
	case int:
		fmt.Printf("x is int type")
	case float64:
		fmt.Printf("x is float64 type")
	case func(int) float64:
		fmt.Printf("x is func(int) type")
	case bool, string:
		fmt.Printf("x is bool or string type")
	default:
		fmt.Printf("Unknown type")
	}
	/*
		when to use switch vs if
		1. if the condition is simple, such as use int, float, char, string, then go switch
		2. if condition is like a range or bool type, then go if
	*/
}
