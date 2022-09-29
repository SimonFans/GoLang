package main

import "fmt"

// introduce the use case for read only & write only channel
// channel only can write
func send(ch chan<- int, exitChan chan bool) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("Send data=%v\n", i)
	}
	close(ch)
	exitChan <- true
}

// channel only can read
func receive(ch <-chan int, exitChan chan bool) {
	for {
		data, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Receive data=%v\n", data)
	}
	exitChan <- true
}

func main() {
	ch := make(chan int, 10)
	exitChan := make(chan bool, 2)
	go send(ch, exitChan)
	go receive(ch, exitChan)
	cnt := 0
	for _ = range exitChan {
		cnt++
		fmt.Printf("%v go routine completes\n", cnt)
		if cnt == 2 {
			break
		}
	}
}
