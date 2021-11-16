package rectangle

type rectangle struct {
	length, breadth float64
}

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}
func (r rectangle) Perimeter() float64 {
	return 2*r.length + 2*r.breadth
}
func New(length, breadth float64) *rectangle {
	return &rectangle{
		length:  length,
		breadth: breadth,
	}
}
