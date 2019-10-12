package main

import (
	"fmt"
	"math"
)

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
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height / 2
}

type Circle struct {
	Diameter float64
}

func (r Circle) Area() float64 {
	return math.Pi * math.Pi * r.Diameter
}

func (r Circle) Perimeter() float64 {
	return math.Pi * r.Diameter
}

func PrintArea(s interface{}) {

	//type assertion

	shape, ok := s.(Shape)

	if !ok {
		fmt.Println("Cannot use type assertion to this type")
		return
	}
	fmt.Println("Shape has area=", shape.Area())
}

func PrintPerimeter(s Shape) {
	fmt.Println("Shape has area=", s.Perimeter())
}
func main() {

	rect := Rect{Width: 20, Height: 30}

	PrintArea(rect)

	cirle := Circle{Diameter: 50}

	PrintArea(cirle)

	triangle := Triangle{Width: 30, Height: 60}

	PrintArea(triangle)
}
