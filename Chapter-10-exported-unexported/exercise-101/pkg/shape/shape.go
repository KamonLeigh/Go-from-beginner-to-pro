package shape

import "fmt"

type Shape interface {
	area() float64
	name() string
	perimeter() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

type Square struct {
	Side float64
}

func (t Triangle) area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) name() string {
	return "triangle"
}

func (t Triangle) perimeter() float64 {
	// assumes equilateral triangle
	return t.Base * 3
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) perimeter() float64 {
	return (r.Length + r.Width) * 2
}

func (r Rectangle) name() string {
	return "rectangle"
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (s Square) perimeter() float64 {
	return s.Side * 4
}

func (s Square) name() string {
	return "square"
}

func PrintShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f and the perimeter: %.2f\n", item.name(), item.area(), item.perimeter())
	}
}
