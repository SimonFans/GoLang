package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Tom\tJack")
	fmt.Println("\"Hello\"")
	// \r: enter but not go to the new line. It will overwrite such as abcde replaces the mochi
	fmt.Println("mochiisadog\rabcde")
	// unsafe.Sizeof(b) returns how many bits it takes. 1 byte = 8 bits
	var b int = 255
	fmt.Printf("b is a type of %T, it takes bytes:%d\n", b, unsafe.Sizeof(b))
	num := 1.5e-2
	fmt.Println(num)
	// Print result: int32 => 97 a
	c1 := 'a'
	fmt.Printf("%T => %v %c\n", c1, c1, c1)
	// c2=Ü
	var c2 int = 220
	fmt.Printf("c2=%c\n", c2)
	// Print raw statement
	str_raw := `package main
import "fmt"`
	fmt.Println(str_raw)
	// "99
	num1 := 99
	num1_to_str := fmt.Sprintf("%d", num1)
	fmt.Printf("%q\n", num1_to_str)
	// "23.456000"
	num2 := 23.456
	num2_to_str := fmt.Sprintf("%f", num2)
	fmt.Printf("%q\n", num2_to_str)
	// "true"
	logic := true
	logic_to_str := fmt.Sprintf("%t", logic)
	fmt.Printf("%q\n", logic_to_str)
	// "h"
	char1 := 'h'
	char_to_str := fmt.Sprintf("%c", char1)
	fmt.Printf("%q\n", char_to_str)
}
