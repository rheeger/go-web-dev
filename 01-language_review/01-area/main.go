package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func printArea(x shape) {
	fmt.Println("My area is", x.area())
}
func main() {
	s1 := square{10}
	c1 := circle{2.3}
	printArea(s1)
	printArea(c1)
}
