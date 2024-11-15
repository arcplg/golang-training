package main

import "fmt"

const Pi = 3.14

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * Pi * c.Radius
}

func PrintValue(v interface{}) {
	fmt.Println("value:", v)
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	var s Shape
	s = Circle{Radius: 5}

	PrintValue(s.Area())
	PrintValue(s.Perimeter())
	PrintArea(s)

	var s1 Shape = Circle{Radius: 10}
	c, ok := s1.(Circle)

	if ok {
		fmt.Println("Circle radius:", c.Radius)
	}
}
