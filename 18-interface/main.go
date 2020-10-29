package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
	perimeter() float64
	// setWidth()
}

// Rect struct
type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * r.width * r.height
}

// func (r Rect) setWidth(width float64) {
// 	r.width = width
// }

// Circle struct
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	// rect := new(Rect)
	// rect.width = 20
	// rect.height = 34

	r := Rect{20., 34.}
	c := Circle{10.}

	// rect.setWidth(23)

	// fmt.Println(r.area())
	// fmt.Println(c.area())

	showArea(r, c)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		// println(a)
		fmt.Println(a)
	}
}
