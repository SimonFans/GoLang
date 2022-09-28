package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// define a channel that can store different types of data, (int, string, map, struct, struct pointer)
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "abd"
	allChan <- Cat{"xiaomao", 5}
	// pop twice, try to get the third struct data
	<-allChan
	<-allChan
	newCat := <-allChan
	// visit Cat property needs to use leixing duanyan. Because newCat.Name returns an interface and interface doesn't have any property.
	c1 := newCat.(Cat)
	fmt.Println(c1.Name)
}
