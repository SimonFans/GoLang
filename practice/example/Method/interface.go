package Method

import (
	"fmt"
	"math"
)

type Circle struct {
	X, Y, Radius float64
}

type Rectangle struct {
	Height, Width float64
}

type Shape interface {
	area() float64
}

func (circle Circle) area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}

func (rectangle Rectangle) area() float64 {
	return rectangle.Height * rectangle.Width
}

func PrintAll(shape []Shape) {
	for _, s := range shape {
		fmt.Printf("%.2f\n", s.area())
	}
}
