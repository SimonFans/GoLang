package main

import "fmt"

// Channel is yin yong (reference) type
func main() {
	// create an int channel with 3 capacity
	intChan := make(chan int, 3)
	fmt.Printf("Channel points to an address=%v\nChannel has its own address=%p\n", intChan, &intChan)
	// write into channel
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 30
	fmt.Printf("Channel length=%v and Channel Capacity=%v\n", len(intChan), cap(intChan))
	// pop from the Channel follows FIFO queue
	num1 := <-intChan
	fmt.Printf("The first number pop from the Channel=%v\n", num1)
	fmt.Printf("Channel length=%v and Channel Capacity=%v\n", len(intChan), cap(intChan))
}
