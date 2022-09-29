package main

import (
	"fmt"
	"time"
)

// print out all prime numbers between 1 to 200

func putNum(ch chan int) {
	for i := 1; i <= 20; i++ {
		ch <- i
	}
}

func isPrime(num int, pch chan int) {
	flag := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		pch <- num
	}
}

func main() {
	intChan := make(chan int, 20)
	primeChan := make(chan int, 20)
	go putNum(intChan)
	// let Main take a break
	time.Sleep(time.Second * 10)
	for {
		select {
		case v := <-intChan:
			fmt.Printf("Read from intChan=%v\n", v)
			isPrime(v, primeChan)
			//time.Sleep(time.Second)
		case v := <-primeChan:
			fmt.Printf("Prime Number=%v\n", v)
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
