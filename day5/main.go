package main

import (
	"fmt"
	task "task/task"
)

func main() {
	cir := task.NewCircle(3.2)
	tri := task.NewTriangle(3, 4, 5, 6)
	rect := task.NewRectangle(2.3, 1.2)
	sq := task.NewSquare(2.5)

	fmt.Println("")
	fmt.Println(cir.Area())
	fmt.Println(cir.Perimeter())

	fmt.Println(tri.Area())
	fmt.Println(tri.Perimeter())

	fmt.Println(rect.Area())
	fmt.Println(rect.Perimeter())

	fmt.Println(sq.Area())
	fmt.Println(sq.Perimeter())

}
