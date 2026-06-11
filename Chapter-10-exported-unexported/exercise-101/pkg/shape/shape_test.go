package shape

import "testing"

func TestTriangleArea(t *testing.T) {
	tri := Triangle{Base: 10, Height: 5}
	if got := tri.area(); got != 25 {
		t.Errorf("Triangle.area() = %v, want 25", got)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tri := Triangle{Base: 10, Height: 5}
	if got := tri.perimeter(); got != 30 {
		t.Errorf("Triangle.perimeter() = %v, want 30", got)
	}
}

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 4, Length: 5}
	if got := r.area(); got != 20 {
		t.Errorf("Rectangle.area() = %v, want 20", got)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{Width: 4, Length: 5}
	if got := r.perimeter(); got != 18 {
		t.Errorf("Rectangle.perimeter() = %v, want 18", got)
	}
}

func TestSquareArea(t *testing.T) {
	s := Square{Side: 6}
	if got := s.area(); got != 36 {
		t.Errorf("Square.area() = %v, want 36", got)
	}
}

func TestSquarePerimeter(t *testing.T) {
	s := Square{Side: 6}
	if got := s.perimeter(); got != 24 {
		t.Errorf("Square.perimeter() = %v, want 24", got)
	}
}
