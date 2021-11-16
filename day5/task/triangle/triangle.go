package triangle

type triangle struct {
	side1, base, side2, height float64
}

func (t triangle) Area() float64 {
	return (((t.base) * (t.height)) / 2)
}
func (t triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.base
}
func New(side1, side2, base, height float64) *triangle {
	return &triangle{
		side1:  side1,
		base:   base,
		side2:  side2,
		height: height,
	}
}
