package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend")
	default:
		fmt.Printf("Today is %v\n", time.Now().Weekday())
	}

	whoami := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm bool type")
		case string:
			fmt.Printf("I'm string type, my name is %v\n", i)
		case int:
			fmt.Println("I'm int type")
		default:
			fmt.Println("I don't know")
		}
	}
	whoami("Ximeng")
}
