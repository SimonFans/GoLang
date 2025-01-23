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
}
