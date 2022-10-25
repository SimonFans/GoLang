package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "aaa"
	fs := reflect.ValueOf(&str1)
	fmt.Println("fs &:", fs)
	fs.Elem().SetString("bbb")
	fmt.Printf("str1 is changed to %v\n", str1)
}
