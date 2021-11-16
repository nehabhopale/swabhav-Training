package task

import "math"

type circle struct {
	radius float64
}
type square struct {
	side float64
}
type rectangle struct {
	length, breadth float64
}
type triangle struct {
	side1, base, side2, height float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (s square) Area() float64 {
	return s.side * s.side
}
func (s square) Perimeter() float64 {
	return 4 * s.side
}
func (r rectangle) Area() float64 {
	return r.length * r.breadth
}
func (r rectangle) Perimeter() float64 {
	return 2*r.length + 2*r.breadth
}
func (t triangle) Area() float64 {
	return (((t.base) * (t.height)) / 2)
}
func (t triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.base
}
func CircleInstance(radius float64) *circle {
	return &circle{
		radius: radius,
	}
}
func RectangleInstance(length, breadth float64) *rectangle {
	return &rectangle{
		length:  length,
		breadth: breadth,
	}
}
func SquareInstance(side float64) *square {
	return &square{
		side: side,
	}
}
func TriangleInstance(side1, side2, base, height float64) *triangle {
	return &triangle{
		side1:  side1,
		base:   base,
		side2:  side2,
		height: height,
	}
}
