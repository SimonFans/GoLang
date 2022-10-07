package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Score [5]float64
	ptr   *int              // pointer
	slice []int             // slice
	map1  map[string]string // map
}

func main() {
	var p1 Person
	num := 5
	fmt.Println(p1)
	p1.Name = "Ximeng"
	p1.Age = 20
	p1.ptr = &num
	p1.slice = make([]int, 5)
	p1.slice[0] = 6
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "val1"
	fmt.Printf("Name=%v, Age=%v, pointer_value=%v, slice=%v, map1=%v\n", p1.Name, p1.Age, *(p1.ptr), p1.slice, p1.map1)
}
