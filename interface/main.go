package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	redius float64
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.redius
}

// Rect 타입에 대한 Shape 인터페이스의 구현
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		fmt.Println(a)
	}
}

func main() {

	r := Rect{width: 10, height: 5}
	c := Circle{redius: 5}

	showArea(r, c)

}
