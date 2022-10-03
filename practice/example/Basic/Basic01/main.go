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
}
