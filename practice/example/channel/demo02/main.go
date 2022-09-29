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

	// Close channel & iterate channel
	intChan2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan2 <- i + 1
	}
	// iterate the channel using for..range.. (note: need to close the channel before doing iteration otherwise it will keep iterating the channel even though it's empty)
	close(intChan2)
	for v := range intChan2 {
		fmt.Printf("v=%v\n", v)
	}
}
