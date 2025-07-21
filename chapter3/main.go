package main

import (
	"fmt"
	"math"
)

type geometry interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}
func (r rect) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

func main() {
	r := rect{width: 5, height: 10}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
