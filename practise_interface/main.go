package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rect{Width: 10, Height: 20}

	PrintArea(rect)

	c := Circle{Diameter: 30}
	PrintArea(c)

}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width  float64
	Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}
func (r Rect) Perimeter() float64 {
	return 0
}

type Circle struct {
	Diameter float64
}

func (r Circle) Area() float64 {
	return r.Diameter * r.Diameter / 4 * math.Pi
}

func PrintArea(s interface{}) {
	if shape, ok := s.(Shape); ok {
		fmt.Println(shape.Area())
	} else {
		fmt.Println("Unmatch interface")
	}

}
