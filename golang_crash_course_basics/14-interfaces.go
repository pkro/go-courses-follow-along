package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	w, h float64
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{1.5}
	rect := Rectangle{2.5, 3.2}

	fmt.Printf("Circle area: %f", getArea(circle))
	fmt.Printf("Rectangle area: %f", getArea(rect))
}
