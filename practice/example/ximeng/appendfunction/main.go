package main

import "fmt"

func main() {
	// s := []int{}
	var s []int
	// []
	fmt.Println(s)
	s = append(s, 0)
	// [0]
	fmt.Println(s)
	s = append(s, 1, 2, 3)
	// [0,1,2,3]
	fmt.Println(s)
	// concatenate two slices
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s1 = append(s1, s2...)
	// The ... unpacks b. Without the dots, the code would attempt to append the slice as a whole, which is invalid.
	// [1 2 3 4]
	fmt.Println("combine s1 and s2 into s1", s1)
	// special case: legal to append string to a byte slice
	special := append([]byte("hello"), "st"...)
	// [104 101 108 108 111 115 116]
	fmt.Println(special)
	g := []byte{}
	g = append(g, 'a')
	// [97]
	fmt.Println(g)
}
