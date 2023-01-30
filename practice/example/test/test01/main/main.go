package main

import (
	"fmt"
	_ "sync"
	"time"
	// "io/ioutil"
)

func printNum() {
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
	}
}

func divideNum() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("divide error: ", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func main() {
	go printNum()
	go divideNum()
	time.Sleep(time.Second * 5)
}

// var wg sync.WaitGroup // only define, no need to assign value

// func read(intChan chan int) {
// 	defer wg.Done()
// 	for v := range intChan {
// 		fmt.Printf("Reading data %v\n", v)
// 		time.Sleep(time.Second)
// 	}
// }

// func write(intChan chan int) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		intChan <- i
// 		fmt.Printf("Writing data %v\n", i)
// 		time.Sleep(time.Second * 3)
// 	}
// 	close(intChan)
// }

// func main() {
// 	intChan := make(chan int, 10)
// 	wg.Add(2)
// 	go write(intChan)
// 	go read(intChan)
// 	wg.Wait()
// }

// inputPath := "/home/xmzhao/mygit/GoLang/practice/example/test/test01/main/test.txt"
// outputPath := "/home/xmzhao/mygit/GoLang/practice/example/test/test01/main/output.txt"
// content, err := ioutil.ReadFile(inputPath)
// if err != nil {
// 	fmt.Println("Read fails: ", err)
// }
// err = ioutil.WriteFile(outputPath, content, 0666)
// if err != nil {
// 	fmt.Println("Write fails: ", err)
// }
