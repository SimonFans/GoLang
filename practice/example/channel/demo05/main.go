package main

import (
	"fmt"
	"time"
)

// use "select" keyword instead of closing the channel. Sometimes it's hard to know when to close the channel.
func main() {
	intChan := make(chan int, 10)
	strChan := make(chan string, 5)
	// read int numbers to intChan
	go func(ch chan int) {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}(intChan)
	// write strings to strChan
	go func(ch chan string) {
		for i := 1; i <= 5; i++ {
			ch <- "Hello" + fmt.Sprintf("%d", i)
		}
	}(strChan)
	// let main wait for 3seconds until write operations to complete
	time.Sleep(time.Second * 3)
	// select
	// label:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("Read from intChan=%v\n", v)
			time.Sleep(time.Second)
		case v := <-strChan:
			fmt.Printf("Read from strChan=%s\n", v)
			time.Sleep(time.Second)
		// default part developer can define what they want to do next
		default:
			fmt.Printf("Cannot read anything...\n")
			// quit the main thread
			return
			// label also works. But label just jump out the for loop. It will continue to run the rest of code in the main thread.
			// break label
		}
	}
}
