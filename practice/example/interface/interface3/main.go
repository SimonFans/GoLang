package main

import "fmt"

// define an interface

type Paobu interface {
	Run()
}

type Cat struct {
}

type Dog struct {
}

// How does cat run
func (c *Cat) Run() {
	fmt.Println("Cat is running...")
}

// How does dog run
func (d *Dog) Run() {
	fmt.Println("Dog is running...")
}

// Empty interface
type T interface {
}

func main() {
	var cat Cat
	var dog Dog
	// 接口指向了实现该接口自定义的变量
	var c1 Paobu = &cat
	c1.Run()
	var d1 Paobu = &dog
	d1.Run()
	// Empty Interface has no method. All types implements empty inferface. Any variables can assigned to an empty interface.
	num := 15
	var t T = num
	fmt.Println(t)
	t = cat
	fmt.Println(t)
}
