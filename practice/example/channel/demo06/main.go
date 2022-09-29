package main

import (
	"fmt"
	"time"
)

// defer + recover to report error when any single go routine has something wrong. This will prevent stopping all go routines to run.

func sayHello() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%v: Hello World!!\n", i)
	}
}

// import defer + recover mechanism to capture any errors from this go routine .
func divideNum() {
	// capture error
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("err=%v\n", err)
		}
	}()
	numerator := 5
	//denumerator := 2
	denumerator := 0
	res := numerator / denumerator
	fmt.Printf("res=%v\n", res)
}

func main() {
	go sayHello()
	go divideNum()
	time.Sleep(time.Second * 8)
	fmt.Printf("Main Program ends...\n")
}
