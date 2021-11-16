package circle

import (
	"math"
)

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func New(radius float64) *circle {
	return &circle{
		radius: radius,
	}
}
