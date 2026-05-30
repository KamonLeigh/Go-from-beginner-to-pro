package main

import "fmt"

type Shape interface {
	Area() float64
	Name() string
	Perimeter() float64
}

type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	width  float64
	length float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{width: 10, length: 20}
	s := square{side: 10}

	printShapeDetails(t, r, s)
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t triangle) Name() string {
	return "triangle"
}

func (t triangle) Perimeter() float64 {
	// assumes equilateral triangle
	return t.base * 3
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func (r rectangle) Perimeter() float64 {
	return (r.length + r.width) * 2
}

func (r rectangle) Name() string {
	return "rectangle"
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Perimeter() float64 {
	return s.side * 4
}

func (s square) Name() string {
	return "square"
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f and the perimeter: %.2f\n", item.Name(), item.Area(), item.Perimeter())
	}
}
