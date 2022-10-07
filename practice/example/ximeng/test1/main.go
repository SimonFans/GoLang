package main

import "fmt"

type Student struct {
	Name string
	Age  int
	ptr  *int
}

func main() {
	var s1 Student
	score := 100
	s1.Name = "Ximeng"
	s1.Age = 10
	s1.ptr = &score
	fmt.Printf("Name=%v, Age=%v, Score=%v", s1.Name, s1.Age, *(s1.ptr))
}
