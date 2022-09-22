package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("The %v parameter type is bool, value is %v\n", index+1, item)
		case float64:
			fmt.Printf("The %v parameter type is float64, value is %v\n", index+1, item)
		case int, int32, int64:
			fmt.Printf("The %v parameter type is int, value is %v\n", index+1, item)
		case string:
			fmt.Printf("The %v parameter type is string, value is %v\n", index+1, item)
		case Student:
			fmt.Printf("The %v parameter type is Student, value is %v\n", index+1, item)
		case *Student:
			fmt.Printf("The %v parameter type is *Student, value is %v\n", index+1, item)
		default:
			fmt.Printf("The %v parameter type is unknown, value is %v\n", index+1, item)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int = 8
	var n4 string = "hello"
	var n5 = make(map[string]int)
	stu1 := Student{}
	stu2 := &Student{}
	fmt.Printf("stu1 %v, stu2 %v\n", stu1, stu2)
	TypeJudge(n1, n2, n3, n4, n5, stu1, stu2)
}
