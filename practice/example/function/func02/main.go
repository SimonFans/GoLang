package main

import "fmt"

var age = getAge()

// run 1
func getAge() int {
	fmt.Println("Get age...")
	return 20
}

// run 2
// init function can complete some intialization works
func init() {
	fmt.Println("Init...")
}

// run 3
func main() {
	fmt.Println("Main...")
	fmt.Println(age)
}
