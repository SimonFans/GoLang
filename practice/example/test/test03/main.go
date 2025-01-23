package main

import "fmt"

type Caculator struct {
	a int
	b int
}

func (c Caculator) add() int {
	return c.a + c.b
}

func (c Caculator) multiply() int {
	return c.a * c.b
}

func adder(a, b int) int {
	return a + b
}

func adder2(a, b int) int {
	return a - b
}

func main() {
	cal := Caculator{1, 3}
	{
		x := cal.add()
		y := cal.multiply()
		fmt.Println(x, y)
	}
	{
		x := cal.add()
		y := cal.multiply()
		fmt.Println(x, y)
	}

	f := []func(a, b int) int{adder, adder2}
	for _, fn := range f {
		f1 := fn(1, 2)
		fmt.Println(f1)
	}
	fmt.Println(f)
}
