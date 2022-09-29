package main

import (
	"fmt"
	"time"
)

// func factorial(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

func writeData(intChan chan int) {
	for i := 1; i <= 20; i++ {
		intChan <- i
		fmt.Printf("Writeing data=%v\n", i)
		time.Sleep(time.Second)
	}
	// close the intChan
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		out, ok := <-intChan
		// no available number in the channel intChan, then stop reading & insert a true to exitChan.
		if !ok {
			break
		}
		fmt.Printf("Reading data=%v\n", out)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	// global intChan: read & write numbers
	intChan := make(chan int, 20)
	// exitChan: notify the main thread that the read operation is done so that the main thread can stop running
	exitChan := make(chan bool, 1)
	// define two goroutine. One is for write 50 numbers to intChan, the other one is to read 50 numbers from intChan.
	go writeData(intChan)
	go readData(intChan, exitChan)
	fmt.Println("Hello while loop!!!")
	// main thread keeps reading the value from exitChan until it see a true. Quit the main thread.
	for {
		fmt.Println("exitChan", exitChan)
		_, ok := <-exitChan
		fmt.Println("Main check ok:", ok)
		if !ok {
			break
		}
	}
	// res := factorial(5)
	// fmt.Printf("Factorial result is %v\n", res)
}
