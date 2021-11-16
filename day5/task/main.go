package main

import (
	"fmt"
	ShapeInterface "task/ShapeInterface"
	circle "task/circle"
	rectangle "task/rectangle"
	square "task/square"
	triangle "task/triangle"
)

func main() {
	cir := circle.New(3.2)
	getDetails(cir)
	tri := triangle.New(3, 4, 5, 6)
	getDetails(tri)
	rect := rectangle.New(2.3, 1.2)
	getDetails(rect)
	sq := square.New(2.5)
	getDetails(sq)
}
func getDetails(shape ShapeInterface.ShapeInterface) {
	fmt.Println("area of a shape-", shape.Area())
	fmt.Println("perimeter of a shape-", shape.Perimeter())
}
